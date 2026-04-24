package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Valor inválido")
		return
	}
	s := 0.0
	num := 1000.0
	sinal := 1.0
	for i := 1.0; i <= float64(n); i++ {
		s += sinal * (num / i)
		num -= 3
		sinal *= -1
	}
	fmt.Println("S =", s)
}