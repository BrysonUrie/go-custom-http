package customhttp

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Uri     string
	Version string
	Headers map[string]string
	Body    string
}

func (req *Request) String() string {
	headerStr := ""
	for key, val := range req.Headers {
		headerStr += fmt.Sprintf("%s: %s\n", key, val)
	}

	return fmt.Sprintf(
		"Method: %s \nUri: %s\nVersion: %s\n%s"+
			"Body: %s",
		req.Method, req.Uri, req.Version, headerStr, req.Body,
	)

}

func createRequest(request string) (*Request, error) {
	lines := strings.Split(request, "\r\n")
	ln1 := strings.Split(lines[0], " ")
	method := ln1[0]
	uri := ln1[1]
	version := ln1[2]

	headers := make(map[string]string)
	last := 1
	for i, header := range lines[1:] {
		last = i
		if header == "" {
			break
		}
		split := strings.Split(header, ":")
		if len(split) < 2 {
			return nil, fmt.Errorf("Header in improper format: %s", header)
		}
		headers[split[0]] = split[1]
	}

	body := strings.Join(lines[last+2:], "\r\n")

	return &Request{
		Method:  method,
		Uri:     uri,
		Version: version,
		Headers: headers,
		Body:    body,
	}, nil
}
