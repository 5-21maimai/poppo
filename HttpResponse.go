package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"os"
	"strconv"
	"unicode/utf8"
)

type HttpResponse struct {
	status string
	header map[string]string
	body   string
}

func NewHttpResponse() *HttpResponse {
	param := new(HttpResponse)
	param.header = map[string]string{}
	return param
}

func (p *HttpResponse) createResponse() string {
	http := "HTTP/1.0 200 OK"
	headerString := ""
	for k, v := range p.header {
		headerString = headerString + "\n" + k + ": " + v
	}
	return http + headerString + "\n\n" + p.body
}

func (p *HttpResponse) addHeader(key string, value string) {
	p.header[key] = value
}

func (p *HttpResponse) addBody(body string) {
	p.body = body
	count := utf8.RuneCountInString(body)
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
		contentType = "image/png"
	} else {
		filename = "index.html"
		p.body = p.createStringBody(filename)
		contentType = "text/html; charset=UTF-8"
	}

	count := utf8.RuneCountInString(p.body)
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

}
