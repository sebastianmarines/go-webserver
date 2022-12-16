package main

import "github.com/sebastianmarines/go-webserver"

func main() {
	webserver.NewWebserver().Start(":8080")
}
