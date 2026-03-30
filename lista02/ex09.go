package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	if x < 10 {
		fmt.Printf("Valor de venda: %f", x*1.7)
	} else if x > 9 && x < 30 {
		fmt.Printf("Valor de venda: %f", x*1.5)
	} else if x > 29 && x < 50 {
		fmt.Printf("Valor de venda: %f", x*1.4)
	} else if x > 49 {
		fmt.Printf("Valor de venda: %f", x*1.3)
	}
}
