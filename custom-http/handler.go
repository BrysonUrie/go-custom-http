package customhttp

import (
	"fmt"
	"syscall"
)

func Listen(portNum int) {
	fd, err := syscall.Socket(
		syscall.AF_INET,
		syscall.SOCK_STREAM,
		syscall.IPPROTO_TCP,
	)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		panic(err)
	}

	addr := syscall.SockaddrInet4{Port: portNum}
	copy(addr.Addr[:], []byte{0, 0, 0, 0})

	if err := syscall.Bind(fd, &addr); err != nil {
		panic(err)
	}

	if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		panic(err)
	}
	fmt.Printf("Listening on :%d\n", portNum)

	for {
		connFd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}

		go handleConnection(connFd)
	}
}

func handleConnection(connFd int) {
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
