package main

import (
	"testing"
)

func TestDummy(t *testing.T) {
	got := Dummy(6)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
