package http

import (
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNewGetRequest(t *testing.T) {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	tcp := l.(*net.TCPListener)
	defer tcp.Close()
	tcp.SetDeadline(time.Now().Add(time.Second))

	go func() {
		_, err := http.Get("http://localhost:8080/example/123")
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := tcp.Accept()
	if err != nil {
		t.Fatal(err)
	}

	req, err := newRequest(conn)
	if err != nil {
		t.Fatalf("expected request created, got error %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("expected method GET, got method %v", req.Method)
	}
	if req.Path != "/example/123" {
		t.Errorf("expected path /example/123, got path %v", req.Path)
	}
	if req.Header["Host"] != "localhost:8080" {
		t.Errorf("expected host localhost:8080, got host %v", req.Header["Host"])
	}
}

func TestNewPostRequest(t *testing.T) {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	tcp := l.(*net.TCPListener)
	defer tcp.Close()
	tcp.SetDeadline(time.Now().Add(time.Second))

	go func() {
		_, err := http.Post("http://localhost:8080/example/456", "text/plain", strings.NewReader("content from request body"))
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := tcp.Accept()
	if err != nil {
		t.Fatal(err)
	}

	req, err := newRequest(conn)
	if err != nil {
		t.Fatalf("expected request created, got error %v", err)
	}

	if req.Method != "POST" {
		t.Errorf("expected method POST, got method %v", req.Method)
	}
	if req.Path != "/example/456" {
		t.Errorf("expected path /example/456, got path %v", req.Path)
	}
	if req.Header["Host"] != "localhost:8080" {
		t.Errorf("expected host localhost:8080, got host %v", req.Header["Host"])
	}

	content := string(req.Content)
	if content != "content from request body" {
		t.Errorf("expected content 'content from request body', got content '%v'", content)
	}
}
