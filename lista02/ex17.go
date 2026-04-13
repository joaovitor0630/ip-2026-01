package main

import "fmt"

func main() {
	var conta, tipo int
	var consumo, total float64
	fmt.Scan(&conta, &tipo, &consumo)

	switch tipo {
	case 1: // Residencial
		total = 5.0 + (consumo * 0.05)
	case 2: // Comercial
		total = 500.0
		if consumo > 800 {
			total += (consumo - 800) * 0.25
		}
	case 3: // Industrial
		total = 800.0
		if consumo > 100 {
			total += (consumo - 100) * 0.04
		}
	}
	fmt.Printf("Conta: %d - Valor a pagar: R$ %.2f\n", conta, total)
}
