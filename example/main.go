package main

import web "github.com/sebastianmarines/go-webserver"

func main() {
	server := web.NewWebserver()
	server.Get("/", func(request web.Request) web.Response {
		return web.HTMLResponse("<h1>Hello World!</h1>", 200, nil)
	})
	server.Start(":8080")
}
