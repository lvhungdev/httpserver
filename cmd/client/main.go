package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			statusCode, content, err := sendReq("GET", "http://localhost:8080/"+strconv.Itoa(i), "")
			if err != nil {
				panic(err)
			}

			fmt.Println(i, statusCode, content)
			wg.Done()
		}()
	}

	statusCode, content, err := sendReq("POST", "http://localhost:8080/", "content from body request")
	if err != nil {
		panic(err)
	}

	fmt.Println("POST", statusCode, content)

	wg.Wait()
}

func sendReq(method string, url string, body string) (int, string, error) {
	var res *http.Response
	var err error

	if method == "GET" {
		res, err = http.Get(url)
	} else if method == "POST" {
		res, err = http.Post(url, "text/plain", strings.NewReader(body))
	}

	if err != nil {
		return 0, "", err
	}

	buffer := make([]byte, 1024)
	n, err := res.Body.Read(buffer)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	content := string(buffer[:n])

	return res.StatusCode, content, nil
}
