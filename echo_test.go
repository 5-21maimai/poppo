package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := sum(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMinus(t *testing.T) {
	actual := minus(10, 20)
	expected := 10
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCreateResponse(t *testing.T) {
	response := NewHttpResponse()
	response.addHeader("Content-Type", "text/html")
	response.addHeader("Server", "maimai")
	response.addHeader("Content-Length", "19")
	response.addBody("<h1>matsui mai</h1>")
	actual := response.createResponse()
	expected := "HTTP/1.0 200 OK\nContent-Type: text/html\nServer: maimai\nContent-Length: 19\n\n<h1>matsui mai</h1>"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
