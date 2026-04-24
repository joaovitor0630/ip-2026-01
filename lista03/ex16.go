package main
import "fmt"
func main() {
	var n, a1, a2 int
	fmt.Scan(&n, &a1, &a2)
	if n < 3 {
		return
	}
	fmt.Print(a1, " ", a2)
	for i := 3; i <= n; i++ {
		var ai int
		if i%2 == 0 {
			ai = a2 - a1
		} else {
			ai = a2 + a1
		}
		fmt.Print(" ", ai)
		a1 = a2
		a2 = ai
	}
	fmt.Println()
}