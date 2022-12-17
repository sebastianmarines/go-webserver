package webserver

import (
	"fmt"
)

func validateHttpStartLine(line string) (method, path, proto string, err error) {
	supportedVerbs := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}
	_, err = fmt.Sscanf(line, "%s %s %s", &method, &path, &proto)
	if err != nil {
		err = fmt.Errorf("invalid http start line: %s", line)
		return "", "", "", err
	}
	validVerb := false
	for _, m := range supportedVerbs {
		if m == method {
			validVerb = true
		}
	}
	if !validVerb {
		err = fmt.Errorf("invalid http verb: %s", method)
		return "", "", "", err
	}
	err = nil
	return
}
