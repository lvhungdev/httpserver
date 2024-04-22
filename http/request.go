package http

import (
	"bytes"
	"errors"
	"io"
	"net"
)

type Request struct {
	Method  string
	Path    string
	Header  map[string]string
	Content []byte
}

func newRequest(conn net.Conn) (Request, error) {
	req := Request{
		Header: make(map[string]string),
	}

	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		return req, err
	}

	data = data[:n]
	reader := bytes.NewReader(data)

	line, err := readLine(reader)
	if err != nil {
		return req, err
	}

	lineSlice := bytes.Split(line, []byte{' '})
	if len(lineSlice) != 3 {
		return req, errors.New("invalid format")
	}

	req.Method = string(lineSlice[0])
	req.Path = string(lineSlice[1])

	for {
		optionalHeader, err := readLine(reader)
		if err != nil {
			return req, err
		}
		if len(optionalHeader) == 0 {
			break
		}

		slice := bytes.Split(optionalHeader, []byte{':', ' '})
		if len(slice) != 2 {
			return req, errors.New("invalid format")
		}

		req.Header[string(slice[0])] = string(slice[1])
	}

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return req, nil
			} else {
				return req, err
			}
		}
		req.Content = append(req.Content, b)
	}
}

func readLine(reader *bytes.Reader) ([]byte, error) {
	data := []byte{}

	for {
		b, err := reader.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			return data, err
		}

		data = append(data, b)

		if len(data) >= 2 && data[len(data)-2] == '\r' {
			return data[:len(data)-2], nil
		}
	}
}
