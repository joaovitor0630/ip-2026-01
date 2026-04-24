package main
import "fmt"
func main() {
	s := 0.0
	num := 1.0
	sinal := 1.0
	for denBase := 15.0; denBase >= 1.0; denBase-- {
		den := denBase * denBase
		s += sinal * (num / den)
		num *= 2
		sinal *= -1
	}
	fmt.Println("S =", s)
}