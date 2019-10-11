package main

import (
	"testing"
)

func TestProductStruct(t *testing.T) {
	// I just want to make sure Product is valid
	product := &Product{
		Name: "warsh cloth",
	}
	if product == nil {
		t.Error("Can't initialize Product")
	}
}
