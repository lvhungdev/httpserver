package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			url := "http://localhost:8080/" + strconv.Itoa(i)
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			buffer := make([]byte, 1024)
			n, err := res.Body.Read(buffer)
			if err != nil && !errors.Is(err, io.EOF) {
				panic(err)
			}

			content := string(buffer[:n])

			fmt.Println(i, res.StatusCode, content)
			wg.Done()
		}()
	}

	wg.Wait()
}
