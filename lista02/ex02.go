package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Print("posivo")
	} else if n < 0 {
		fmt.Print("negativo")
	} else {
		fmt.Print("nulo")
	}

}
