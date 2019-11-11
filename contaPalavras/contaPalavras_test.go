package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
