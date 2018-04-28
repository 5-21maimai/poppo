package main

import (
	"strings"
	"testing"
)

func TestNewHttpResponse(t *testing.T) {
	response := NewHttpResponse()

	actual := response.status
	expected := ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = response.body
	expected = ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual2 := len(response.header)
	expected2 := 0
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAddBody(t *testing.T) {
	response := NewHttpResponse()
	response.addBody("hoge")
	actual := response.body
	expected := "hoge"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAddHeader(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("hoge", "fuga")
	actual := response.header["hoge"]
	expected := "fuga"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateStatusLine(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBody("<h1>matsui mai</h1>")

	responseHeader := strings.Split(response.createResponse(), "\n")
	actual := responseHeader[0]
	expected := "HTTP/1.0 200 OK"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateHeader(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBody("<h1>matsui mai</h1>")

	responseHeader := response.createResponse()
	actual := strings.Contains(responseHeader, "Content-Type: text/html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = strings.Contains(responseHeader, "Content-Length: 19")
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBody(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBody("<h1>matsui mai</h1>")

	responseHeader := response.createResponse()
	actual := strings.Contains(responseHeader, "<h1>matsui mai</h1>")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyHtmlDefault(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBodyHtml("/")

	responseString := response.createResponse()
	actual := strings.Contains(responseString, "index.html")
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyHtmlHello(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBodyHtml("/hello")

	responseString := response.createResponse()
	actual := strings.Contains(responseString, "hello.html")
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyHtmlHoge(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addBodyHtml("/hoge")

	responseString := response.createResponse()
	actual := strings.Contains(responseString, "index.html")
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
