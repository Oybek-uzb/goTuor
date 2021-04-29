package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	emptyResult := hello("")

	if emptyResult != "Hello, Dude" {
		t.Errorf("hello(\"\") failed. Expected %v, got %v", "Hello, Dude", emptyResult)
	} else {
		t.Logf("hello(\"\") success. Expected %v, got %v", "Hello, Dude", emptyResult)
	}

	result := hello("Mike")

	if result != "Hello, Mike" {
		t.Errorf("hello(\"Mike\") failed. Expected %v, got %v", "Hello, Mike", result)
	} else {
		t.Logf("hello(\"Mike\") success. Expected %v, got %v", "Hello, Mike", result)
	}
}
