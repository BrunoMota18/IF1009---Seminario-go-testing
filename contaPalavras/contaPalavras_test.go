/*
https://golang.org/cmd/go/#hdr-Testing_flags
go test -v: verbose output
go test -bench ".": benchmark
go test -run <F>
go test -cover
go test -coverprofile="cover.txt"
go tool cover -html="cover.txt" -o cover.html
*/
package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	got := Dummy(6)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestContaPalavras(t *testing.T) {
	assertEqualMaps := func(t *testing.T, got, want map[string]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("checa palavras de uma frase", func(t *testing.T) {
		gotDicionario, gotErr := contaPalavras("Hasta la vista baby vista la baby hasta la")
		wantDicionario := map[string]int{
			"Hasta": 1,
			"baby":  2,
			"la":    3,
			"vista": 2,
			"hasta": 1,
		}

		if gotErr != nil {
			t.Errorf("got %q want %q", gotErr, "nil")
		}
		assertEqualMaps(t, gotDicionario, wantDicionario)
	})

	t.Run("checa se acusa erro quando não há frase", func(t *testing.T) {
		gotDicionario, gotErr := contaPalavras("")
		wantErr := errors.New("o dicionario está vazio; certifique se digitou alguma frase")
		assert.Equal(t, 0, len(gotDicionario))
		assert.Equal(t, wantErr, gotErr)
	})
}

func BenchmarkContaPalavras(b *testing.B) {
	for i := 0; i < b.N; i++ {
		contaPalavras(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
					   sed do eiusmod tempor incididunt ut labore et dolore magna 
					   aliqua. Ac turpis egestas sed tempus urna et pharetra. Duis 
					   at consectetur lorem donec massa sapien faucibus. Viverra ipsum 
					   nunc aliquet bibendum enim. Dui accumsan sit amet nulla facilisi 
					   morbi tempus iaculis. Blandit volutpat maecenas volutpat blandit 
					   aliquam etiam erat. Augue ut lectus arcu bibendum at. Pharetra diam 
					   sit amet nisl suscipit adipiscing bibendum. Pharetra diam sit amet 
					   nisl suscipit adipiscing bibendum est ultricies. Dolor sit amet consectetur 
					   adipiscing elit pellentesque.`)
	}
}
