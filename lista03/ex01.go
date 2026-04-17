package main

import "fmt"

func main() {

	var x, n int

	fmt.Scan(&x, &n)
	fmt.Print(pet(x, n))

}

func pet(n, y int) int {

	if y == 0 {
		return 1
	}

	return n * pet(n, y-1)
}
