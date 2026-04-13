package main

import "fmt"

func main() {
	var preco float64
	var dia, cat int
	fmt.Print("Preço Normal, Dia (1-7), Categoria (1-Comum, 2-Lançamento): ")
	fmt.Scan(&preco, &dia, &cat)

	final := preco
	// Desconto por dia [cite: 68]
	if dia == 2 || dia == 3 || dia == 5 {
		final *= 0.60
	}
	// Acréscimo por categoria [cite: 71]
	if cat == 2 {
		final *= 1.15
	}
	fmt.Printf("Preço final: R$ %.2f\n", final)
}
