package webserver

import (
	"strings"
)

type RouteNode struct {
	Path     string
	Handlers map[string]HandlerFunc
	Child    []*RouteNode
}

type RouteTree struct {
	Root *RouteNode
}

type HandlerFunc func(Request) Response

func (s *Server) addRoute(method string, path string, handler HandlerFunc) {
	segments := strings.Split(path, "/")
	node := s.routes.Root
	if node == nil {
		node = &RouteNode{}
		s.routes.Root = node
	}
	// traverse tree
	var lastSegment string
	for i, segment := range segments {
		// Check if it is root path "/" and if it has a handler
		if lastSegment == "" && segment == "" && i != 0 {
			if node.Handlers == nil {
				node.Handlers = make(map[string]HandlerFunc)
			}
			node.Handlers[method] = handler
			return
		}
		lastSegment = segment
		// check if segment exists
		if segment == "" {
			continue
		}
		// check if segment is a parameter
		found := false
		for _, child := range node.Child {
			// check if segment is a parameter
			if child.Path == segment {
				node = child
				found = true
				break
			}
		}
		// if segment is not a parameter, create a new node
		if !found {
			newNode := RouteNode{
				Path:     segment,
				Handlers: make(map[string]HandlerFunc),
				Child:    make([]*RouteNode, 0),
			}
			// register handler if it is the last segment
			if segment == segments[len(segments)-1] {
				newNode.Handlers[method] = handler
			}
			node.Child = append(node.Child, &newNode)
			node = &newNode
		}
	}
}

func (s *Server) handleRoute(r Request) Response {
	segments := strings.Split(r.Path, "/")
	node := s.routes.Root
	if node == nil {
		return NotFoundResponse()
	}
	// traverse tree
	var lastSegment string
	for i, segment := range segments {
		// Check if it is root path "/" and if it has a handler
		if lastSegment == "" && segment == "" && i != 0 {
			if handler, ok := node.Handlers[r.Method]; ok {
				return handler(r)
			}
			return NotFoundResponse()
		}
		lastSegment = segment
		// check if segment exists
		if segment == "" {
			continue
		}
		// check if segment is a parameter
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
			return NotFoundResponse()
		}
	}
	// check if handler exists
	handler, ok := node.Handlers[r.Method]
	if ok {
		return handler(r)
	}
	return NotFoundResponse()
}

func (s *Server) Get(path string, handler func(Request) Response) {
	s.addRoute("GET", path, handler)
}

func (s *Server) Post(path string, handler func(Request) Response) {
	s.addRoute("POST", path, handler)
}

func (s *Server) Put(path string, handler func(Request) Response) {
	s.addRoute("PUT", path, handler)
}

func (s *Server) Delete(path string, handler func(Request) Response) {
	s.addRoute("DELETE", path, handler)
}

func (s *Server) Patch(path string, handler func(Request) Response) {
	s.addRoute("PATCH", path, handler)
}

func (s *Server) Head(path string, handler func(Request) Response) {
	s.addRoute("HEAD", path, handler)
}

func (s *Server) Options(path string, handler func(Request) Response) {
	s.addRoute("OPTIONS", path, handler)
}

func (s *Server) Connect(path string, handler func(Request) Response) {
	s.addRoute("CONNECT", path, handler)
}

func (s *Server) Trace(path string, handler func(Request) Response) {
	s.addRoute("TRACE", path, handler)
}
