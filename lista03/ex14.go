package main
import "fmt"
func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for i := n1; i <= n2; i++ {
		if i < 2 {
			continue
		}
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}