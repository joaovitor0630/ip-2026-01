package main
import "fmt"
func main() {
	s := 0.0
	num := 100.0
	fat := 1.0
	for i := 0; i < 20; i++ {
		if i > 0 {
			fat *= float64(i)
		}
		s += num / fat
		num--
	}
	fmt.Println("S =", s)
}