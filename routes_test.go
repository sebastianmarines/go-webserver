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
