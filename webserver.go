package webserver

import (
	"bufio"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type Server struct {
	port   string
	routes RouteTree
}

func NewWebserver() *Server {
	server := Server{}
	server.routes = RouteTree{}
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
			request.PathParams = make(map[string]string)
			request.QueryParams = make(map[string]string)

			startLine, err := reader.ReadString('\n')
			method, path, proto, err := validateHttpStartLine(startLine)

			if err != nil {
				response := TextResponse("<h1>400 Bad Request</h1>", 400, nil)
				_, err := c.Write(response.build())
				if err != nil {
					return
				}
				return
			}

			request.Method = method
			request.Path = path
			request.Proto = proto

			log.Printf("Request: %s", startLine)

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

			u, err := url.Parse(request.Path)
			if err != nil {
				log.Fatal(err)
			}
			request.Path = u.Path
			queryParams := u.Query()
			for k, v := range queryParams {
				request.QueryParams[k] = v[0]
			}

			response := s.handleRoute(request)

			_, err = c.Write(response.build())
			if err != nil {
				return
			}

		}(conn)
	}
}
