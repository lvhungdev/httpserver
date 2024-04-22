package http

import "strconv"

type Response struct {
	statusCode  int
	statusText  string
	contentType string
	content     []byte
}

func NewResponse(statusCode int, statusText string, contentType string, content []byte) Response {
	return Response{
		statusCode:  statusCode,
		statusText:  statusText,
		contentType: contentType,
		content:     content,
	}
}

func (r Response) Encode() []byte {
	data := []byte{}

	data = append(data, []byte("HTTP/1.1 "+strconv.Itoa(r.statusCode)+" "+r.statusText+"\r\n")...)
	data = append(data, []byte("Content-Type: "+r.contentType+"\r\n")...)
	data = append(data, []byte("Content-Length: "+strconv.Itoa(len(r.content))+"\r\n\r\n")...)

	if r.content != nil {
		data = append(data, r.content...)
	}

	return data
}
