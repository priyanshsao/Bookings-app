package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {

	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("the type is not http.Handler, it is of type %T", v)
	}
}

func TestSessionLoad(t *testing.T) {

	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("the type is not http.Handler, it is of type %T", v)
	}
}
