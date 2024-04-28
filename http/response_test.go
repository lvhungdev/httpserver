package http

import "testing"

func TestResponseEncode(t *testing.T) {
    expected :="HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 19\r\n\r\ncontent from server"

	res := NewResponse(200, "OK", "text/plain", []byte("content from server"))
	got := string(res.Encode())

	if expected != got {
		t.Fatalf("expected '%v', got '%v'", expected, got)
	}
}
