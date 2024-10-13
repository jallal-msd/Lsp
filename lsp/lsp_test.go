package lsp_test

import (
	"rpc/lsp"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	msg := "Content-Length: 16\r\n\r\n {\"Testing\":true}"
	excpect := lsp.Encode(EncodingExample{Testing: true})
	if excpect != msg {
		t.Fatalf("Excpected %s,  but got %s", excpect, msg)
	}
}

func TestDecode(t *testing.T) {
	msg := "Content-Length: 16\r\n\r\n {\"Method\":\"hi\"}"
	method, actual, err := lsp.Decode([]byte(msg))
	if err != nil {
		t.Fatal(err)
	}

	if len(actual) != 16 {
		t.Fatalf("we excpected %d , and we got %d", 16, actual)
	}

	if method != "hi" {
		t.Fatalf("we excpected hi , and we got %s", method)

	}

}
