package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	if n > 99 || n < 1000 {
		n = n % 100
		fmt.Print((n/10)*10, "\n")
	}

}
