package main

import (
	"fmt"
)

func main() {
	var n, b int

	fmt.Scan(&n)
	fmt.Scan(&b)

	if n%b == 0 {
		fmt.Print("é divisível")
	} else {
		fmt.Print("não é")
	}

}
