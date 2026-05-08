package main

import "fmt"

func main() {

	var numeros [10]int
	var maior []int

	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
		if numeros[i] > 50 {
			maior = append(maior, numeros[i])
			maior = append(maior, i+1)
		}
	}
	fmt.Println("Os números maiores que 50 são:")
	for i := 0; i < len(maior); i++ {
		fmt.Printf("%d na posição %d\n", maior[i], maior[i+1])
		i++
	}
}
