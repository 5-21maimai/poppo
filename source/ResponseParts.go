package main

import (
	"strconv"
	"strings"
)

type ResponseParts struct {
	contentTypes map[string]string
}

func NewResponseParts() *ResponseParts {
	param := new(ResponseParts)
	param.contentTypes = map[string]string{
		"html": "text/html",
		"css":  "text/css",
		"png":  "image/png",
		"gif":  "image/gif",
		"js":   "text/javascript",
	}
	return param
}

// bodyからcontentlengthを計算
func makeContentLength(body string) string {
	count := len(body)
	return strconv.Itoa(count)
}

// filenameからcontenttypeを設定
func makeContentType(filename string) string {
	a := strings.Split(filename, ".")
	kakuchoshi := a[len(a)-1]
	renponseParts := NewResponseParts()
	contentType := renponseParts.contentTypes[kakuchoshi]

	return contentType + "; charset=UTF-8"
}
