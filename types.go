package webserver

import "fmt"

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func (r Response) String() []byte {
	response := ""
	response += fmt.Sprintf("HTTP/1.1 %d\r;", r.StatusCode)
	for k, v := range r.Headers {
		response += fmt.Sprintf("%s: %s\r;", k, v)
	}
	response += "\r;"
	response += r.Body
	return []byte(response)
}
