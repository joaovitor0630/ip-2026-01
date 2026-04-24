package main
import "fmt"
func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	a, b := n1, n2
	for b != 0 {
		a, b = b, a%b
	}
	mmc := (n1 * n2) / a
	fmt.Println("MMC:", mmc)
}