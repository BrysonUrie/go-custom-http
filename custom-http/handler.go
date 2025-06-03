package customhttp

import (
	"fmt"
	"syscall"
)

func HandleConnection(connFd int) {
	defer syscall.Close(connFd)

	buf := make([]byte, 1024)
	n, err := syscall.Read(connFd, buf)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	request := string(buf[:n])

	req, _ := createRequest(request)
	fmt.Println(req)

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello, world!"

	_, err = syscall.Write(connFd, []byte(response))
	if err != nil {
		fmt.Println("Write failed:", err)
	}

}
