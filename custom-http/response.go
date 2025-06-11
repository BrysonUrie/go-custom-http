package customhttp

import (
	"fmt"
)

type Response struct {
	Status int16
	Body   string
}

var statusText = map[int16]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
}

func CreateResponse(status int16, body string) (*Response, error) {
	if status < 100 || status >= 600 {
		return nil, fmt.Errorf("Invalid error code")
	}
	return &Response{
		Status: status,
		Body:   body,
	}, nil
}

func (response *Response) serializeResponse() string {
	reason, ok := statusText[response.Status]
	if !ok {
		reason = "Unknown Status"
	}

	return fmt.Sprintf("HTTP/1.1 %d %s\r\n", response.Status, reason) +
		"Content-Type: text/plain\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(response.Body)) +
		"\r\n" +
		response.Body
}

func CreateNotFoundRes() *Response {
	res, _ := CreateResponse(404, "No resource found\r\n")
	return res
}

func CreateInternalErrorRes() *Response {
	res, _ := CreateResponse(500, "Error processing request\r\n")
	return res
}
