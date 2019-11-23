package main

import (
	"testing"
)

func TestDummy(t *testing.T) {
	t.Run("deve retornar true se x é maior que 5", func(t *testing.T) {
		got := Dummy(6)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	/*
		t.Run("deve retornar false se x é menor que 5", func(t *testing.T) {
			got := Dummy(2)
			want := false

			if got != want {
				t.Errorf("got %t want %t", got, want)
			}
		})
	*/
}
