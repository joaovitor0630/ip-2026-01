package main

import "fmt"

func main() {
	var raio, altura float64
	fmt.Scan(&raio, &altura)

	const pi = 3.14159
	areaCirculo := pi * raio * raio
	areaLateral := 2 * pi * raio * altura
	areaTotal := 2*areaCirculo + areaLateral

	custo := areaTotal * 100.0

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
