package main
import "fmt"
func main() {
	h := 0.0
	num := 1.0
	for den := 1.0; den <= 50; den++ {
		h += num / den
		num += 2
	}
	fmt.Println("H =", h)
}