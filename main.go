package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Method string
	Path   string
	Proto  string
	Header []Header
	Body   string
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	log.Print("Running on port 8080\n")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			fmt.Printf("New connection from %s\n", c.RemoteAddr())
			reader := bufio.NewReader(c)

			startLine, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Request: %s", startLine)

			request := Request{}
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
				request.Header = append(request.Header, header)
			}
			// Check if we have a Content-Length header
			for _, header := range request.Header {
				if header.Key == "Content-Length" {
					// Read the body
					size, err := strconv.Atoi(header.Value)
					if err != nil {
						log.Fatal(err)
					}
					body := make([]byte, size)
					_, err = reader.Read(body)
					if err != nil {
						return
					}
					request.Body = string(body)
				}
			}

			fmt.Printf("Closing connection from %s\n", c.RemoteAddr())

			response := "HTTP/1.1 200 OK\r\n\r\nHello World!"

			// Send the response back to the client
			_, err = conn.Write([]byte(response))
			if err != nil {
				return
			}

			// Print the request
			fmt.Printf("Request: %s\n", prettyPrint(request))

		}(conn)
	}
}
