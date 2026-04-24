package main
import (
	"fmt"
	"math"
)
func main() {
	s := 0.0
	denBase := 1.0
	sinal := 1.0
	for i := 0; i < 51; i++ {
		den := denBase * denBase * denBase
		s += sinal * (1.0 / den)
		denBase += 2
		sinal *= -1
	}
	pi := math.Cbrt(s * 32.0)
	fmt.Println("Pi =", pi)
}