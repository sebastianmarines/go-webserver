package webserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	port   string
	routes map[string]Route
}

type Route struct {
	Path    string
	Handler func(Request) Response
}

func NewWebserver() *Server {
	server := Server{}
	server.routes = make(map[string]Route)
	return &server
}

func (s *Server) Start(a string) {

	log.Printf("Running on %s\n", a)
	ln, err := net.Listen("tcp", a)
	if err != nil {
		log.Fatal(err)
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {

		}
	}(ln)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer func(c net.Conn) {
				err := c.Close()
				if err != nil {

				}
			}(c)

			log.Printf("New connection from %s\n", c.RemoteAddr())
			reader := bufio.NewReader(c)

			request := Request{}
			request.Headers = make(map[string]string)

			startLine, err := reader.ReadString('\n')
			method, path, proto, err := validateHttpStartLine(startLine)

			request.Method = method
			request.Path = path
			request.Proto = proto

			if err != nil {
				response := Response{
					StatusCode: 400,
					Headers:    make(map[string]string),
					Body:       "Bad Request",
				}
				response.Headers["Content-Length"] = strconv.Itoa(len(response.Body))
				response.Headers["Content-Type"] = "text/plain"
				response.Headers["Connection"] = "close"
				_, err := c.Write(response.String())
				if err != nil {
					return
				}
				return
			}

			fmt.Printf("Request: %s", startLine)

			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}

				// End of headers
				if line == "\r\n" {
					break
				}
				header := Header{}
				parts := strings.Split(line, ":")
				header.Key = parts[0]
				if len(parts) > 2 {
					header.Value = strings.Join(parts[1:], ":")
				} else {
					header.Value = parts[1]
				}
				header.Value = strings.Trim(header.Value, " ")
				// Remove "\r\n"
				header.Value = header.Value[:len(header.Value)-2]
				request.Headers[header.Key] = header.Value
			}

			// Check if we have a Content-Length header
			i, ok := request.Headers["Content-Length"]
			if ok {
				// Read the body
				length, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
				}
				body := make([]byte, length)
				_, err = reader.Read(body)
				if err != nil {
					log.Fatal(err)
				}
				request.Body = string(body)
			}

			response := s.handleRoute(request)

			response.Headers["Content-Length"] = strconv.Itoa(len(response.Body))
			// If content type is not set, set it to text/plain
			_, ok = response.Headers["Content-Type"]
			if !ok {
				response.Headers["Content-Type"] = "text/plain"
			}
			response.Headers["Connection"] = "close"
			_, err = c.Write(response.String())
			if err != nil {
				return
			}

		}(conn)
	}
}

func (s *Server) AddRoute(r string, h func(Request) Response) {
	s.routes[r] = Route{Path: r, Handler: h}
}

func (s *Server) handleRoute(r Request) Response {
	i, ok := s.routes[r.Path]
	if ok {
		return i.Handler(r)
	}
	return Response{
		StatusCode: 404,
		Headers:    make(map[string]string),
		Body:       "Not Found",
	}
}
