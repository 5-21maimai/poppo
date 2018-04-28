package main

import (
	"bufio"
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

func (p *HttpResponse) addBodyHtml(path string) {
	str := ""
	filename := "index.html"
	if path == "/hello" {
		filename = "hello.html"
	}

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

	p.body = str

	count := utf8.RuneCountInString(str)
	p.addHeader("Content-Length", strconv.Itoa(count))
}
