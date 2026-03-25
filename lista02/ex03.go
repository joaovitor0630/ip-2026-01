package main

import "fmt"

func main() {

	var n1, n2 int

	fmt.Scan(&n1, &n2)

	if n1+n2 > 20 {
		fmt.Print(n1 + n2 + 8)
	} else {
		fmt.Print((n1 + n2) - 5)
	}

}
