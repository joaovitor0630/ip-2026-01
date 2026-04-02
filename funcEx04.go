package main

import "fmt"

func main() {

	var valor_fatorial int64
	fmt.Scan(&valor_fatorial)

	fmt.Print(fatorial(valor_fatorial), "\n")

}

func fatorial(numero int64) int64 {

	for i := numero - 1; i != 0; i-- {
		numero *= i
	}
	return numero

}
