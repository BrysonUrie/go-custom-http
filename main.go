package main

import (
	"fmt"

	"github.com/brysonurie/go-http/custom-http"
)

func test(handler *customhttp.Request) string {
	fmt.Println("AHHHHH")
	return ""
}

func main() {
	router, _ := customhttp.CreateRouter()
	router.RegisterPath("/hello", test)
	router.Listen(8080)
}
