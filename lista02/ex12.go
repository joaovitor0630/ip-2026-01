package main

import (
	"fmt"
)

func main() {
	var idade int

	fmt.Print("Digite a idade da pessoa: ")

	var classificacao string

	if idade <= 2 {
		classificacao = "Recém-nascido"
	} else if idade <= 11 {
		classificacao = "Criança"
	} else if idade <= 19 {
		classificacao = "Adolescente"
	} else if idade <= 55 {
		classificacao = "Adulto"
	} else {
		classificacao = "Idoso"
	}

	// Exibe o resultado formatado
	fmt.Printf("Classificação: %s\n", classificacao)
}
