package webserver

import "testing"

func TestValidateHttpStartLine(t *testing.T) {
	method, path, proto, err := validateHttpStartLine("GET / HTTP/1.1")
	if err != nil {
		t.Error("Error validating http start line")
	}
	if method != "GET" {
		t.Error("Invalid method")
	}
	if path != "/" {
		t.Error("Invalid path")
	}
	if proto != "HTTP/1.1" {
		t.Error("Invalid protocol")
	}

	// Test an invalid verb
	method, path, proto, err = validateHttpStartLine("GETT / HTTP/1.1")
	if err == nil {
		t.Error("Error validating http start line")
	}

	// Test an invalid start line
	method, path, proto, err = validateHttpStartLine("GETT /")
	if err == nil {
		t.Error("Error validating http start line")
	}
}
