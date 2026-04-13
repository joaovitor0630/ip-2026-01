package main

import "fmt"

func main() {
	var preco float64
	var cod int
	fmt.Scan(&preco, &cod)

	switch cod {
	case 1:
		fmt.Println(preco * 0.90)
	case 2:
		fmt.Println(preco * 0.95)
	case 3:
		fmt.Println(preco)
	case 4:
		fmt.Println(preco * 1.10)
	}
}
