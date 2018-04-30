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
