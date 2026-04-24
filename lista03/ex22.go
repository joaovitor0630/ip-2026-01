package main
import "fmt"
func main() {
	s := 0.0
	den := 1.0
	for i := 37.0; i >= 1.0; i-- {
		s += (i * (i + 1)) / den
		den++
	}
	fmt.Println("S =", s)
}