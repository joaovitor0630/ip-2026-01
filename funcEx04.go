package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	fmt.Print(fatorial(a))

}

func fatorial(a int) int {

	for i := a - 1; i != 0; i-- {
		a *= i
	}
	return a

}
