package main

import "example.com/http/http"

func main() {
	s := http.NewServer()

	s.Handle("POST", "/", func(r *http.Request) (http.Response, error) {
		return http.NewResponse(200, "OK", "text/plain", r.Content), nil
	})

	s.Handle("GET", "/1", func(r *http.Request) (http.Response, error) {
		return http.NewResponse(200, "OK", "text/plain", []byte("everything is good so far")), nil
	})

	s.Handle("GET", "/7", func(r *http.Request) (http.Response, error) {
		return http.NewResponse(200, "OK", "text/plain", []byte("everything is good so far")), nil
	})

	err := s.ListenAndServe(":8080")

	if err != nil {
		panic(err)
	}
}
