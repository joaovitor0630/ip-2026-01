package main
import "fmt"
func main() {
	for a := 0.0; a <= 6.3001; a += 0.1 {
		a3 := a * a * a
		a5 := a3 * a * a
		a7 := a5 * a * a
		senA := a - (a3 / 6.0) + (a5 / 120.0) - (a7 / 5040.0)
		fmt.Printf("Sen(%.1f) = %f\n", a, senA)
	}
}