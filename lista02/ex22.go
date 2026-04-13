package main

import "fmt"

func main() {
	var mat, hEx int
	fmt.Scan(&mat, &hEx)

	bruto := (3 * 788.0) + (float64(hEx) * 10.0)
	desc := 0.0
	if bruto > 1500 {
		desc += bruto * 0.12
	}
	if bruto > 2000 {
		desc += bruto * 0.20
	}

	fmt.Printf("Matrícula: %d, Líquido: R$ %.2f\n", mat, bruto-desc)
}
