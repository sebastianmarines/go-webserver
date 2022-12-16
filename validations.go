package webserver

import (
	"fmt"
)

func validateHttpStartLine(line string) (method, path, proto string, err error) {
	_, err = fmt.Sscanf(line, "%s %s %s", &method, &path, &proto)
	if err != nil {
		err = fmt.Errorf("invalid http start line: %s", line)
		return "", "", "", err
	}
	err = nil
	return
}
