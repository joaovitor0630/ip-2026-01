package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string
	fmt.Scan(&conta, &consumo, &tipo)

	var valor float64

	switch tipo {
	case "R":
		valor = 5.00 + (consumo * 0.05)
	case "C":
		valor = 500.00
		if consumo > 80 {
			valor += (consumo - 80) * 0.25
		}
	case "I":
		valor = 800.00
		if consumo > 100 {
			valor += (consumo - 100) * 0.04
		}
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
