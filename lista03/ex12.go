package main
import "fmt"
func main() {
	var x float64
	fmt.Scan(&x)
	s := x
	sinal := -1.0
	fat := 1.0
	for i := 1; i < 20; i++ {
		fat *= float64(i)
		s += sinal * (x / fat)
		sinal *= -1
	}
	fmt.Println("S =", s)
}