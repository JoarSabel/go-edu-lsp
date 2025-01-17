package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic (err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}

	contentLenghtBytes := header[len("Content-Length: "):]
	contentLenth, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLenth], &baseMessage); err != nil{
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLenth], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error){
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLenghtBytes := header[len("Content-Length: "):]
	contentLenth, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLenth {
		return 0, nil, nil
	}

	totalLenght := len(header) + 4 + contentLenth
	return totalLenght, data[:totalLenght], nil

}
