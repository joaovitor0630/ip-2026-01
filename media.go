package main

import "fmt"

func main() {

	var (
		n int
		notas, soma float64
	)

	fmt.Print("Quantas notas quer calcular: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("digite a nota %d = ", i+1)
		fmt.Scan(&notas)

		soma += notas
	}

	fmt.Printf("a média global da turma é %.2f", soma/float64(n))
}
