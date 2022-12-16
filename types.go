package webserver

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
