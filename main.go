package main

import (
	"fmt"

	"github.com/brysonurie/go-http/custom-http"
)

func test(handler *customhttp.Request) (*customhttp.Response, error) {
	fmt.Println("Hit test handler")
	res, _ := customhttp.CreateResponse(200, "Message Received\r\n")
	return res, nil
}

func testError(handler *customhttp.Request) (*customhttp.Response, error) {
	return nil, fmt.Errorf("This is a test error")
}

func main() {
	router, _ := customhttp.CreateRouter()
	router.RegisterPath("/hello", test)
	router.RegisterPath("/helloError", testError)
	router.Listen(8080)
}
