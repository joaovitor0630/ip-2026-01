package main
import "fmt"
func main() {
	var total uint64 = 0
	var atual uint64 = 1
	for i := 1; i <= 64; i++ {
		total += atual
		atual *= 2
	}
	fmt.Println("Total de grãos:", total)
}