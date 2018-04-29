package main

import (
	"log"
	"os"
	"strconv"
	"strings"
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

// HttpResponse.bodyにコンテンツを追加する（ファイルを読み込む）
func (p *HttpResponse) addBodyPartsFromFile(path string) {
	filename := ""

	switch path {
	case "/hello":
		filename = "hello.html"
		p.body = p.createStringBody(filename)
	case "/":
		filename = "index.html"
		p.body = p.createStringBody(filename)
	case "/methodNotAllowed":
		filename = "methodNotAllowed.html"
		p.body = p.createStringBody(filename)
	case "/hogehoge":
		defer func() {
			if err := recover(); err != nil {
				log.Printf("パニック！！！: %v", err)
				filename = "internalServerError.html"
				p.body = p.createStringBody(filename)
				p.status = "500 Internal Server Error"
			}
		}()
		panic("Occured panic!")
	default:
		filename = strings.Trim(path, "/")
		if _, err := os.Stat(filename); err != nil {
			// not found
			filename = "notFound.html"
			p.body = p.createStringBody(filename)
			p.status = "404 Not Found"
		} else {
			if strings.Contains(filename, ".png") || strings.Contains(filename, ".gif") {
				p.body = p.createBinaryBody(filename)
			} else {
				p.body = p.createStringBody(filename)
			}
		}

	}

	// bodyからcontentlengthを計算してheaderに追加
	p.addHeaderParts("Content-Length", makeContentLength(p.body))

	// filenameからcontenttypeを設定してheaderに追加
	p.addHeaderParts("Content-Type", makeContentType(filename))
}
