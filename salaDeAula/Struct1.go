package main

import "fmt"

type imc struct {
	altura float64
	nome   string
	peso   float64
	ideal  float64
}

func main() {

	var pessoas []imc
	var pessoa imc

	for i := 0; i != -1; i++ {

		fmt.Print("DIGITE SEU NOME E SUA ALTURA RESPECTIVAMENTE SEPARADOS POR ESPAÇO: ")

		fmt.Scan(&pessoa.nome, &pessoa.altura, &pessoa.peso)
		pessoas = append(pessoas, pessoa)
		if pessoa.nome == "FIM" {
			break
		}
	}
	for i := 0; i < len(pessoas); i++ {
		fmt.Print(pessoas[i])
	}
}
