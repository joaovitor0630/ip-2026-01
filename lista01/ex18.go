package main

import "fmt"

func main() {
	var a1, razao, n int
	fmt.Scan(&a1, &razao, &n)

	soma := 0
	termoAtual := a1

	for i := 0; i < n; i++ {
		soma += termoAtual
		termoAtual += razao
	}

	fmt.Printf("%d\n", soma)
}
