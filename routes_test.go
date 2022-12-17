package webserver

import "testing"

// test add route
func TestAddRoute(t *testing.T) {
	server := NewWebserver()
	server.Get("/hello/:greeting", func(request Request) Response {
		return HTMLResponse("<h1>Hello World!</h1>", 200, nil)
	})
	// Traverse the tree to see if the route was added
	node := server.routes.Root
	if node == nil {
		t.Error("Root node is nil")
	}
	for _, segment := range []string{"hello", ":greeting"} {
		if segment == "" {
			continue
		}
		found := false
		for _, child := range node.Child {
			// check if segment is a parameter
			if child.Path == segment {
				node = child
				found = true
				break
			}
		}
		if !found {
			t.Error("Route not added")
		}
	}
}

func TestAddRootRoute(t *testing.T) {
	server := NewWebserver()
	server.Get("/", func(request Request) Response {
		return HTMLResponse("<h1>Hello World!</h1>", 200, nil)
	})
	// Traverse the tree to see if the route was added
	node := server.routes.Root
	if node == nil {
		t.Error("Root node is nil")
	}
	if node.Handlers["GET"] == nil {
		t.Error("Route not added")
	}
}

func TestPathParameters(t *testing.T) {
	server := NewWebserver()
	server.Get("/hello/:greeting", func(request Request) Response {
		if request.PathParams["greeting"] != "world" {
			t.Error("Path parameter not set")
		}
		return HTMLResponse("<h1>Hello World!</h1>", 200, nil)
	})
	request := Request{
		Method: "GET",
		Path:   "/hello/world",
	}
	response := server.handleRoute(request)
	if response.StatusCode != 200 {
		t.Error("Route not found")
	}
}
