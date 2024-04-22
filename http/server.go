package http

import (
	"fmt"
	"net"
)

type Server struct {
	router Router
}

func NewServer() *Server {
	return &Server{
		router: newRouter(),
	}
}

func (s *Server) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection %v", err.Error())
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			req, err := newRequest(conn)
			if err != nil {
				res := NewResponse(500, "Internal Server Error", "text/plain", []byte(err.Error()))
				conn.Write(res.Encode())
				return
			}

			handler := s.router.getHandler(req.Method, req.Path)
			if handler == nil {
				res := NewResponse(404, "Not Found", "text/plain", []byte{})
				conn.Write(res.Encode())
				return
			}

			res, err := handler(&req)
			if err != nil {
				res := NewResponse(500, "Internal Server Error", "text/plain", []byte{})
				conn.Write(res.Encode())
				return
			}

			conn.Write(res.Encode())
		}(conn)
	}
}

func (s *Server) Handle(method string, path string, handler HandleFunc) {
	s.router.addHandler(method, path, handler)
}
