package customhttp

import (
	"fmt"
	"syscall"
)

type Router struct {
	Paths map[string]*Path
}

func CreateRouter() (*Router, error) {
	return &Router{
		Paths: map[string]*Path{},
	}, nil
}

func (router *Router) RegisterPath(path string, handler func(*Request) (*Response, error)) error {
	pathObj := &Path{
		PurePath: path,
		Handler:  handler,
	}
	router.Paths[path] = pathObj
	return nil
}

func (router *Router) Listen(portNum int) {
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

		go router.handleConnection(connFd)
	}
}

func (router *Router) handleConnection(connFd int) {
	defer syscall.Close(connFd)

	buf := make([]byte, 1024)
	n, err := syscall.Read(connFd, buf)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	request := string(buf[:n])

	req, _ := createRequest(request)

	handler, err := router.getHandler(req.Uri)
	var res *Response = nil

	if err != nil {
		res = CreateNotFoundRes()
	} else {
		res, err = handler.Handler(req)
		if err != nil {
			res = CreateInternalErrorRes()
		}
	}

	_, err = syscall.Write(connFd, []byte(res.serializeResponse()))
	if err != nil {
		fmt.Println("Write failed:", err)
	}
}

func (router *Router) getHandler(goal string) (*Path, error) {
	for path, handler := range router.Paths {
		if path == goal {
			return handler, nil
		}
	}
	return nil, fmt.Errorf("No handler found")
}
