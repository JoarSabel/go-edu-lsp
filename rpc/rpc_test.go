package rpc_test

import (
	"lsp-test-project/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
 func TestDecode(t *testing.T){
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected lenght %d but got %d", 15, contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected %s but got %s", "hi", method)
	}
}
