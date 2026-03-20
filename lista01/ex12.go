package main

import "fmt"

func main() {
	var horas int
	fmt.Scan(&horas)

	blocosDe3 := horas / 3
	horasRestantes := horas % 3

	valor := float64(blocosDe3*10 + horasRestantes*5)
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
