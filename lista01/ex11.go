package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
