package webserver

import "strconv"

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
	MediaType  string
}

func (r Response) build() []byte {
	statusCode := strconv.Itoa(r.StatusCode)
	statusMessage, ok := statusCodes[statusCode]
	if !ok {
		statusMessage = "Unknown"
	}
	response := "HTTP/1.1 " + statusCode + " " + statusMessage + "\r\n"
	for k, v := range r.Headers {
		response += k + ": " + v + "\r\n"
	}
	response += "Content-Length: " + strconv.Itoa(len(r.Body)) + "\r\n"
	response += "Content-Type: " + r.MediaType + "\r\n"
	response += "\r\n"

	response += r.Body
	return []byte(response)
}

func NewResponse(content string, status int, headers map[string]string, mediaType string) Response {
	return Response{
		StatusCode: status,
		Headers:    headers,
		Body:       content,
		MediaType:  mediaType,
	}
}

func JSONResponse(content string, status int, headers map[string]string) Response {
	if headers == nil {
		headers = make(map[string]string)
	}
	if status == 0 {
		status = 200
	}
	return NewResponse(content, status, headers, "application/json")
}

func HTMLResponse(content string, status int, headers map[string]string) Response {
	if headers == nil {
		headers = make(map[string]string)
	}
	if status == 0 {
		status = 200
	}
	return NewResponse(content, status, headers, "text/html")
}

func TextResponse(content string, status int, headers map[string]string) Response {
	if headers == nil {
		headers = make(map[string]string)
	}
	if status == 0 {
		status = 200
	}
	return NewResponse(content, status, headers, "text/plain")
}

func NotFoundResponse() Response {
	return TextResponse("<h1>404 Not Found</h1>", 404, nil)
}
