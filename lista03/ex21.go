package main
import "fmt"
func main() {
	var b, n int
	fmt.Scan(&b, &n)
	res := 1
	for i := 0; i < n; i++ {
		res *= b
	}
	fmt.Println("Resultado:", res)
}