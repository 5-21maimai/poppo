package main

import (
	"testing"
)

func TestNewHttpRequest(t *testing.T) {
	request := NewHttpRequest()

	actual := request.method
	expected := "GET"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = "/"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadHeaderDefaultGet(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("GET / HTTP/1.0")

	actual := request.method
	expected := "GET"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = "/"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadHeaderHogeGet(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("GET /hoge HTTP/1.0")

	actual := request.method
	expected := "GET"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = "/hoge"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
