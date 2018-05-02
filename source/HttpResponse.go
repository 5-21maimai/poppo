package main

import (
	"strconv"
)

type HttpResponse struct {
	status string
	header map[string]string
	body   string
}

func NewHttpResponse() *HttpResponse {
	param := new(HttpResponse)
	param.header = map[string]string{}
	param.status = "200 OK"
	param.addHeaderParts("Server", "poppo")
	return param
}

// レスポンスを作る大元のメソッド
func (p *HttpResponse) createResponse(method string) string {
	switch method {
	case "GET":
		return p.createGetResponse()
	case "HEAD":
		return p.createHeadResponse()
	default:
		return p.create405Response()
	}
}

// 405用のレスポンス
func (p *HttpResponse) create405Response() string {
	p.addBodyPartsFromFile("/methodNotAllowed")
	p.status = "405 Method Not Allowed"
	http := "HTTP/1.0" + " " + p.status
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

// 500用のレスポンス
func (p *HttpResponse) create500Response() string {
	p.addBodyPartsFromFile("/internalServerError")
	p.status = "500 Internal Server Error"
	http := "HTTP/1.0" + " " + p.status
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

// GETのレスポンス
func (p *HttpResponse) createGetResponse() string {
	http := "HTTP/1.0" + " " + p.status
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

// HEADのレスポンス
func (p *HttpResponse) createHeadResponse() string {
	http := "HTTP/1.0 200 OK"
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n"
}

// HttpResponse.headerのmapに項目を追加する
func (p *HttpResponse) addHeaderParts(key string, value string) {
	p.header[key] = value
}

// HttpResponse.bodyにコンテンツを追加する（文字列）
func (p *HttpResponse) addBodyParts(body string) {
	p.body = body
	count := len(body)
	p.addHeaderParts("Content-Length", strconv.Itoa(count))
	p.addHeaderParts("Content-Type", "text/html; charset=UTF-8")

}
