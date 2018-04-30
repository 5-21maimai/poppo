package main

import (
	"strings"
	"testing"
)

func TestNewHttpResponse(t *testing.T) {
	response := NewHttpResponse()

	actual := response.status
	expected := "200 OK"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = response.body
	expected = ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual2 := len(response.header)
	expected2 := 1
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAddHeader(t *testing.T) {
	response := NewHttpResponse()
	response.addHeaderParts("hoge", "fuga")
	actual := response.header["hoge"]
	expected := "fuga"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateStatusLine200(t *testing.T) {
	response := NewHttpResponse()

	responseHeader := strings.Split(response.createResponse("GET"), "\n")
	actual := responseHeader[0]
	expected := "HTTP/1.0 200 OK"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	responseHeader = strings.Split(response.createResponse("HEAD"), "\n")
	actual = responseHeader[0]
	expected = "HTTP/1.0 200 OK"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateStatusLine405(t *testing.T) {
	response := NewHttpResponse()

	responseHeader := strings.Split(response.createResponse("hoge"), "\n")
	actual := responseHeader[0]
	expected := "HTTP/1.0 405 Method Not Allowed"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateHeaderFromString(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyParts("<h1>matsui mai</h1>")

	responseHeader := response.createResponse("GET")
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

func TestCreateBodyFromIndexHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/")

	actual := strings.Contains(response.body, "index.html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyFromHelloHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/hello")

	actual := strings.Contains(response.body, "hello.html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyFromNotAllowedHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/methodNotAllowed")

	actual := strings.Contains(response.body, "methodNotAllowed.html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyFromNotFoundHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/fugafuga")

	actual := strings.Contains(response.body, "notFound.html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyFromPathHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/sample.js")

	actual := strings.Contains(response.body, "sample.js")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateBodyFromPanic(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/permissionDenied.html")

	actual := strings.Contains(response.body, "internalServerError.html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateHeaderFromNotFoundHtml(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/fugafuga")

	actual := strings.Contains(response.status, "404 Not Found")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateHeaderFromPanic(t *testing.T) {
	response := NewHttpResponse()
	response.addBodyPartsFromFile("/permissionDenied.html")

	actual := strings.Contains(response.status, "500 Internal Server Error")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateResponseGET(t *testing.T) {
	response := NewHttpResponse()
	response.body = "<h1>hello world</h1>"
	response.header["Content-Length"] = "20"
	response.header["Content-Type"] = "text/html"
	str := response.createResponse("GET")

	actual := strings.Contains(str, "Content-Type: text/html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = strings.Contains(str, "Content-Length: 20")
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = strings.Contains(str, "\n\n<h1>hello world</h1>")
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateResponseHEAD(t *testing.T) {
	response := NewHttpResponse()
	response.body = "<h1>hello world</h1>"
	response.header["Content-Length"] = "20"
	response.header["Content-Type"] = "text/html"
	str := response.createResponse("HEAD")

	actual := strings.Contains(str, "Content-Type: text/html")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = strings.Contains(str, "Content-Length: 20")
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = strings.Contains(str, "\n\n<h1>hello world</h1>")
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
