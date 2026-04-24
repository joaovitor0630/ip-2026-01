package main
import "fmt"
func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	mult := 0
	absN2 := n2
	if n2 < 0 {
		absN2 = -n2
	}
	for i := 0; i < absN2; i++ {
		if n2 < 0 {
			mult -= n1
		} else {
			mult += n1
		}
	}
	fmt.Println("Multiplicação:", mult)
}