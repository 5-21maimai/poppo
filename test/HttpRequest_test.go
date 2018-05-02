package main

import (
	"testing"
)

func TestNewHttpRequest(t *testing.T) {
	request := NewHttpRequest()

	actual := request.method
	expected := ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadHeaderDefaultGET(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("GET / HTTP/1.1")

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

func TestReadHeaderHogeGET(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("GET /hoge HTTP/1.1")

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

func TestReadHeaderDefaultHEAD(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("HEAD / HTTP/1.1")

	actual := request.method
	expected := "HEAD"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = "/"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadHeaderHogeHEAD(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("HEAD /hoge HTTP/1.1")

	actual := request.method
	expected := "HEAD"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = "/hoge"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadHeaderInvalid(t *testing.T) {
	request := NewHttpRequest()
	request.readHeader("hogehoge")

	actual := request.method
	expected := ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = request.path
	expected = ""
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
