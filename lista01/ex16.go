package main

import "fmt"

func main() {
	var salario float64
	fmt.Scan(&salario)

	var reajuste float64
	if salario <= 300.00 {
		reajuste = salario * 1.50
	} else {
		reajuste = salario * 1.30
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", reajuste)
}
