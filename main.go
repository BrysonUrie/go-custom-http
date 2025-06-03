package main

import (
	"fmt"
	"github.com/brysonurie/go-http/custom-http"
	"syscall"
)

func main() {
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

	addr := syscall.SockaddrInet4{Port: 8080}
	copy(addr.Addr[:], []byte{0, 0, 0, 0})

	if err := syscall.Bind(fd, &addr); err != nil {
		panic(err)
	}

	if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		panic(err)
	}
	fmt.Println("Listening on :8080")

	for {
		connFd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}

		go customhttp.HandleConnection(connFd)
	}

}
