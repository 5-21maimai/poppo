package main

import (
	"testing"
)

func TestMakeContentLength(t *testing.T) {
	actual := makeContentLength("hogehoge")

	expected := "8"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMakeContentTypeHtml(t *testing.T) {
	actual := makeContentType("hoge.html")

	expected := "text/html; charset=UTF-8"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMakeContentTypeCss(t *testing.T) {
	actual := makeContentType("hoge.css")

	expected := "text/css; charset=UTF-8"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMakeContentTypeJs(t *testing.T) {
	actual := makeContentType("hoge.js")

	expected := "text/javascript; charset=UTF-8"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
