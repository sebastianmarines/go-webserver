package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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

			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}

				if line == "\r\n" {
					fmt.Printf("End of headers\n")
					break
				}
				fmt.Printf("Header: %q\n", line)
			}

			fmt.Printf("Closing connection from %s\n", c.RemoteAddr())

			response := "HTTP/1.1 200 OK\r\n\r\nHello World!"

			// Send the response back to the client
			conn.Write([]byte(response))

			//for {
			//	msg, err := reader.ReadString('\n')
			//	if err == io.EOF {
			//		break
			//	}
			//	if err != nil {
			//		fmt.Printf("Connection from %s closed\n", c.RemoteAddr())
			//		return
			//	}
			//	fmt.Printf("Message: %s\n", msg)
			//	i, err := c.Write([]byte("Hello from server"))
			//	if err != nil {
			//		fmt.Printf("Connection from %s closed\n", c.RemoteAddr())
			//		return
			//	}
			//	fmt.Printf("Wrote %d bytes\n", i)
			//}
		}(conn)
	}
}
