package main

import "fmt"

func main() {
	var (
		menor, inter, maior, aux int
	)

	fmt.Scan(&menor, &inter, &maior)

	if menor > inter {
		aux = menor
		menor = inter
		inter = aux
	}
	if menor > maior {
		aux = menor
		menor = maior
		maior = aux
	}
	if inter > maior {
		aux = inter
		inter = maior
		maior = aux
	}

	fmt.Printf("%d  %d  %d\n", menor, inter, maior)

}
