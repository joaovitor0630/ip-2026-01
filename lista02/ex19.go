package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var r, h float64
	fmt.Print("1-Cone, 2-Cilindro, 3-Esfera: ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Scan(&r, &h)
		vol := (math.Pi * r * r * h) / 3
		area := math.Pi * r * math.Sqrt(r*r+h*h)
		fmt.Printf("V: %.2f, A: %.2f\n", vol, area)
	case 2:
		fmt.Scan(&r, &h)
		vol := math.Pi * r * r * h
		area := 2 * math.Pi * r * h
		fmt.Printf("V: %.2f, A: %.2f\n", vol, area)
	case 3:
		fmt.Scan(&r)
		vol := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		area := 4 * math.Pi * r * r
		fmt.Printf("V: %.2f, A: %.2f\n", vol, area)
	}
}
