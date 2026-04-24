package main
import (
	"fmt"
	"math"
)
func main() {
	for r := 0.0; r <= 20.001; r += 0.5 {
		vol := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		fmt.Printf("Raio: %.1f, Volume: %f\n", r, vol)
	}
}