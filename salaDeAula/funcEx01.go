package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Print(max(a, b, c))

}

func max(a, b, c int) int {
	var maior int = a
	if b > maior {
		maior = b
	}
	if c > maior {
		maior = c
	}
	fmt.Print("\n")
	return maior
}
