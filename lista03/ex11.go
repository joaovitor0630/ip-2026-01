package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Valores inválidos!")
		return
	}
	fat := 1
	for i := 2; i <= n; i++ {
		fat *= i
	}
	fmt.Println("Fatorial:", fat)
}