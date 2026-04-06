package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		a, b, c, delta, x1, x2 float64
	)
	fmt.Print("digite os valores de a b e c:\n")
	fmt.Scan(&a, &b, &c)

	delta = b*b - 4*a*c

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		x1 = -b / (2 * a)
		fmt.Println("RAIZ ÚNICA")
		fmt.Printf("x = %.2f\n", x1)
	} else {
		x1 = (-b + math.Sqrt(delta)) / (2 * a)
		x2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("RAÍZES DISTINTAS")
		fmt.Printf("x1 = %.2f\n", x1)
		fmt.Printf("x2 = %.2f\n", x2)
	}

}
