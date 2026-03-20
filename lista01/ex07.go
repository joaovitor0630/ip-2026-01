package main

import "fmt"

func main() {
	var f, polegadas float64
	fmt.Scan(&f, &polegadas)

	c := (5.0*f - 160.0) / 9.0
	mm := polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", c)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", mm)
}
