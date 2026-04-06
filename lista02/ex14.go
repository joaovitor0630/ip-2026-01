package main

import "fmt"

func main() {
	var precoBase float64
	var escolha string

	fmt.Print("Digite o preço de fábrica: ")
	fmt.Scan(&precoBase)

	precoFinal := precoBase

	// Pergunta cada opcional individualmente
	fmt.Print("Ar condicionado (R$ 1750,00)? (s/n): ")
	fmt.Scan(&escolha)
	if escolha == "s" || escolha == "S" {
		precoFinal += 1750
	}

	fmt.Print("Pintura metálica (R$ 800,00)? (s/n): ")
	fmt.Scan(&escolha)
	if escolha == "s" || escolha == "S" {
		precoFinal += 800
	}

	fmt.Print("Vidro elétrico (R$ 1200,00)? (s/n): ")
	fmt.Scan(&escolha)
	if escolha == "s" || escolha == "S" {
		precoFinal += 1200
	}

	fmt.Print("Direção hidráulica (R$ 2000,00)? (s/n): ")
	fmt.Scan(&escolha)
	if escolha == "s" || escolha == "S" {
		precoFinal += 2000
	}

	// Resultado final
	fmt.Printf("\nPreço de fábrica: R$ %.2f\n", precoBase)
	fmt.Printf("Preço final do carro: R$ %.2f\n", precoFinal)
}
