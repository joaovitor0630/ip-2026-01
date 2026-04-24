package main
import (
	"fmt"
	"math"
)
func main() {
	var x float64
	fmt.Scan(&x)
	cosCalc := 1.0
	sinal := -1.0
	fat := 1.0
	num := 1.0
	for i := 1; i < 20; i++ {
		fat *= float64(2*i) * float64(2*i-1)
		num *= x * x
		cosCalc += sinal * (num / fat)
		sinal *= -1
	}
	cosReal := math.Cos(x)
	fmt.Printf("a) Cosseno calculado: %f\n", cosCalc)
	fmt.Printf("b) Diferença: %f\n", math.Abs(cosCalc-cosReal))
	fmt.Printf("c) Calculado: %f, Função math: %f\n", cosCalc, cosReal)
}