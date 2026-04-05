package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	if n == 2 {
		fmt.Print("ERRO - divisor 0")
	} else {
		fmt.Print(8 / (2 - n))
	}
}
