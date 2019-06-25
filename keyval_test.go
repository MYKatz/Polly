package main

import (
	"testing"
)

func TestSetAndGet(t *testing.T) {
	kv := NewKeyVal() //look at keyval.go
	err := kv.Set("testkey", "42")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	v, err := kv.Get("testkey")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if v != "42" {
		t.Errorf("Error: expected %s but got %s", "42", v)
	}
}
