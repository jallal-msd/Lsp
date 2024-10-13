package lsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func Encode(msg any) string {
	content, err := json.Marshal(msg)

	if err != nil {

		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMethod struct {
	Method string `json:"method"`
}

func Decode(incomingMessage []byte) (string, []byte, error) {

	header, content, found := bytes.Cut(incomingMessage, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("Not found")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		panic(err)
	}
	var baseMethod BaseMethod
	if err := json.Unmarshal(content[:contentLength], &baseMethod); err != nil {
		return "", nil, err
	}

	return baseMethod.Method, content[:contentLength], nil

}

func Split(data []byte, _ bool) (advanced int, token []byte, err error) {

	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if contentLength < len(content) {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength
	return totalLength, header[:totalLength], nil
}
