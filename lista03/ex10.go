package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n < 3 {
		return
	}
	a, b := 0, 1
	fmt.Print(a, " ", b)
	for i := 2; i < n; i++ {
		c := a + b
		fmt.Print(" ", c)
		a = b
		b = c
	}
	fmt.Println()
}