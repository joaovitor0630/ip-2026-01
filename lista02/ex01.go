package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print("É par\n")
	} else {
		fmt.Print("É ímpar\n")
	}
}
