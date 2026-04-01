package main

import "fmt"

func main() {

	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Printf("%.2f", media(a, b, c))

}

func media(a, b, c float64) float64 {
	var media float64 = (a + b + c) / 3
	return media
}
