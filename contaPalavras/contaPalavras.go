package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func contaPalavras(frase string) (map[string]int, error) {
	dicionario := make(map[string]int)
	for _, palavra := range strings.Fields(frase) {
		dicionario[palavra]++
	}
	//fmt.Println(len(dicionario))
	if len(dicionario) == 0 {
		return nil, errors.New("o dicionario está vazio; certifique se digitou alguma frase")
	}
	return dicionario, nil
}

func main() {
	fmt.Println(contaPalavras("Hasta la vista baby"))
	fmt.Println(contaPalavras("Hasta la vista baby vista la baby hasta la"))
	if _, err := contaPalavras("   "); err != nil {
		log.Fatal("Erro!: ", err)
	}
}
