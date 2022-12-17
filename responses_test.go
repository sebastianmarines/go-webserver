package webserver

import (
	"testing"
)

func TestNewResponse(t *testing.T) {
	response := NewResponse("Hello World!", 200, nil, "text/plain")
	if response.Headers == nil {
		t.Error("Headers is nil")
	}
	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if response.Body != "Hello World!" {
		t.Error("Invalid body")
	}
}

func TestJSONResponse(t *testing.T) {
	response := JSONResponse("Hello World!", 200, nil)
	if response.Headers == nil {
		t.Error("Headers is nil")
	}
	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if response.Body != "Hello World!" {
		t.Error("Invalid body")
	}
	if response.MediaType != "application/json" {
		t.Error("Invalid media type")
	}
}

func TestHTMLResponse(t *testing.T) {
	response := HTMLResponse("Hello World!", 200, nil)
	if response.Headers == nil {
		t.Error("Headers is nil")
	}
	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if response.Body != "Hello World!" {
		t.Error("Invalid body")
	}
	if response.MediaType != "text/html" {
		t.Error("Invalid media type")
	}
}

func TestTextResponse(t *testing.T) {
	response := TextResponse("Hello World!", 200, nil)
	if response.Headers == nil {
		t.Error("Headers is nil")
	}
	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if response.Body != "Hello World!" {
		t.Error("Invalid body")
	}
	if response.MediaType != "text/plain" {
		t.Error("Invalid media type")
	}
}

func TestNotFoundResponse(t *testing.T) {
	response := NotFoundResponse()
	if response.Headers == nil {
		t.Error("Headers is nil")
	}
	if response.StatusCode != 404 {
		t.Error("Invalid status code")
	}
	if response.Body != "Not Found" {
		t.Error("Invalid body")
	}
	if response.MediaType != "text/plain" {
		t.Error("Invalid media type")
	}
}
