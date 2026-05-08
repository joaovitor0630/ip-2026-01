package main

import "fmt"

func main() {
	var a [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("A[%d]: ", i)
		fmt.Scan(&a[i])
	}

	fmt.Println("\nElementos repetidos:")
	for i := 0; i < 10; i++ {
		count := 1
		for j := i + 1; j < 10; j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count > 1 {
			fmt.Printf("Valor %d aparece %d vezes\n", a[i], count)
		}
	}
}
