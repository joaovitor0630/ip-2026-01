package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x > 20 && x < 90 {
		fmt.Println("O número digitado está entre 20 e 90")
	} else {
		fmt.Println("O número digitado não está entre 20 e 90")

	}
}
