package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta float64
	fmt.Scan(&altura, &aresta)

	areaBase := (3.0 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2.0
	volume := (1.0 / 3.0) * areaBase * altura

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}
