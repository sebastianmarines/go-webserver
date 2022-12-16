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
	port string
}

func NewWebserver() *Server {
	return &Server{}
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

			startLine, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Request: %s", startLine)

			request := Request{}
			request.Headers = make(map[string]string)
			_, err = fmt.Sscanf(startLine, "%s %s %s", &request.Method, &request.Path, &request.Proto)
			if err != nil {
				// TODO: Handle error
				log.Fatal(err)
			}

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

			fmt.Printf("Closing connection from %s\n", c.RemoteAddr())

			response := "HTTP/1.1 200 OK\r\n\r\nHello World!"

			// Send the response back to the client
			_, err = conn.Write([]byte(response))
			if err != nil {
				return
			}

		}(conn)
	}
}
