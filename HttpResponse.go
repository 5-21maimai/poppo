package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/gif"
	"image/png"
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
	return param
}

func (p *HttpResponse) createResponse(method string) string {
	switch method {
	case "GET":
		return p.createGetResponse()
	case "HEAD":
		return p.createHeadResponse()
	default:
		return p.createInvalidResponse()
	}
}

func (p *HttpResponse) createInvalidResponse() string {
	p.status = "405 Method Not Allowed"
	p.addBodyFile("/methodNotAllowed")
	http := "HTTP/1.0" + " " + p.status
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

func (p *HttpResponse) createGetResponse() string {
	http := "HTTP/1.0" + " " + p.status
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

func (p *HttpResponse) createHeadResponse() string {
	http := "HTTP/1.0 200 OK"
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n"
}

func (p *HttpResponse) addHeader(key string, value string) {
	p.header[key] = value
}

func (p *HttpResponse) addBody(body string) {
	p.body = body
	count := len(body)
	p.addHeader("Content-Length", strconv.Itoa(count))
}

func (p *HttpResponse) addBodyFile(path string) {
	filename := ""
	contentType := ""

	if path == "/hello" {
		filename = "hello.html"
		p.body = p.createStringBody(filename)
		contentType = "text/html; charset=UTF-8"
	} else if path == "/sample.css" {
		filename = "sample.css"
		p.body = p.createStringBody(filename)
		contentType = "text/css"
	} else if path == "/panda.png" {
		filename = "panda.png"
		p.body = p.createBinaryBody(filename)
		contentType = "image/png; charset=UTF-8"
	} else if path == "/twinsparrot.gif" {
		filename = "twinsparrot.gif"
		p.body = p.createBinaryBody(filename)
		contentType = "image/gif; charset=UTF-8"
	} else if path == "/sample.js" {
		filename = "sample.js"
		p.body = p.createStringBody(filename)
		contentType = "text/javascript; charset=UTF-8"
	} else if path == "/" {
		filename = "index.html"
		p.body = p.createStringBody(filename)
		contentType = "text/html; charset=UTF-8"
	} else if path == "/methodNotAllowed" {
		filename = "methodNotAllowed.html"
		p.body = p.createStringBody(filename)
		contentType = "text/html; charset=UTF-8"
	} else {
		filename = "notFound.html"
		p.body = p.createStringBody(filename)
		contentType = "text/html; charset=UTF-8"
		p.status = "404 Not Found"
	}

	count := len(p.body)
	p.addHeader("Content-Length", strconv.Itoa(count))

	p.addHeader("Content-Type", contentType)
}

func (p *HttpResponse) createStringBody(filename string) string {
	str := ""
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		str = str + sc.Text() + "\n"
	}

	return str

}

func (p *HttpResponse) createBinaryBody(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
	}

	if strings.Contains(filename, ".png") {
		img, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			fmt.Println(err)
		}
		imageBytes := buffer.Bytes()
		return string(imageBytes)
	} else if strings.Contains(filename, ".gif") {
		img, err := gif.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		buffer := new(bytes.Buffer)
		if err := gif.Encode(buffer, img, nil); err != nil {
			fmt.Println(err)
		}

		imageBytes := buffer.Bytes()
		return string(imageBytes)
	}

	return "hoge"

}
