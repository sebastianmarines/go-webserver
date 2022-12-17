package webserver

func (s *Server) addRoute(method string, path string, handler func(Request) Response) {
	route := Route{
		Path:    path,
		Handler: handler,
	}
	s.routes[method+path] = route
}

func (s *Server) handleRoute(r Request) Response {
	route, ok := s.routes[r.Method+r.Path]
	if ok {
		return route.Handler(r)
	}
	return Response{
		StatusCode: 404,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body: "<h1>404 Not Found</h1>",
	}
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
