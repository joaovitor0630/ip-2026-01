package main

import "fmt"

func main() {
	var (
		escolha_destino int
		ida_volta       int
	)

	fmt.Print("ESCOLHA O DESTINO DA SUA VIAJEM\n1 - REGIÃO NORTE\n2 - REGIÃO NORDESTE\n3 - REGIÃO CENTRO-OESTE\n4 - REGIÃO SUL\n")
	fmt.Scan("   \n")

	fmt.Scan("digite o número correspondente a sua região")
	fmt.Scan(&escolha_destino)

	fmt.Printf("Destino escolhido para a região %d , ", escolha_destino)

	fmt.Print("A passagem será:\n1 - IDA\n2 - IDA EVOLTA\n")
	fmt.Scan(&ida_volta)

	if ida_volta == 1 {
		if escolha_destino == 1 {
			fmt.Printf("O valor da passagem será R$ 500.")
		} else if escolha_destino == 2 || escolha_destino == 3 {
			fmt.Printf("O valor da passagem será R$ 350.")
		} else if escolha_destino == 4 {
			fmt.Printf("O valor da passagem será R$ 300.")
		}
	}
	if ida_volta == 2 {
		if escolha_destino == 1 {
			fmt.Printf("O valor da passagem será R$ 900.")
		} else if escolha_destino == 2 {
			fmt.Printf("O valor da passagem será R$ 650.")
		} else if escolha_destino == 3 {
			fmt.Printf("O valor da passagem será R$ 600.")
		} else if escolha_destino == 4 {
			fmt.Printf("O valor da passagem será R$ 550.")
		}
	}
}
