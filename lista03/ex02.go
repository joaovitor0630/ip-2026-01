package main
import "fmt"
func main() {
	soma, count := 0, 0
	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			count++
		}
	}
	fmt.Printf("Soma: %d\nMédia: %.2f\n", soma, float64(soma)/float64(count))
}