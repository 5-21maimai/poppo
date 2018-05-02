package main

import "strings"

type HttpRequest struct {
	method string
	path   string
}

func NewHttpRequest() *HttpRequest {
	param := new(HttpRequest)
	param.method = ""
	param.path = ""
	return param
}

func (p *HttpRequest) readHeader(header string) {
	headerArray := strings.Split(header, " ")
	if len(headerArray) > 1 {
		p.method = headerArray[0]
		p.path = headerArray[1]
	}
}
