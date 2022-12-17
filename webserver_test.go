package webserver

import "testing"

func TestQueryParams(t *testing.T) {
	server := NewWebserver()
	server.Get("/hello", func(request Request) Response {
		if request.QueryParams["greeting"] != "world" {
			t.Error("Query parameter not set")
		}
		return HTMLResponse("<h1>Hello World!</h1>", 200, nil)
	})
	request := Request{
		Method: "GET",
		Path:   "/hello?greeting=world",
	}

	server.handleRoute(request)
}
