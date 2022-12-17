package main

import web "github.com/sebastianmarines/go-webserver"

func main() {
	server := web.NewWebserver()
	server.Get("/", func(request web.Request) web.Response {
		return web.Response{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
			Body: "<h1>Hello world!</h1>",
		}
	})
	server.Start(":8080")
}
