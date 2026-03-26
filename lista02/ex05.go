package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	if n%5 == 0 {
		fmt.Print("é divisível")
	} else {
		fmt.Print("não é")
	}

}
