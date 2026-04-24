package main
import "fmt"
func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n <= 0 {
			break
		}
		isPerfect := false
		for i := 1; i*i <= n; i++ {
			if i*i == n {
				isPerfect = true
				break
			}
		}
		if isPerfect {
			fmt.Println("É quadrado perfeito")
		} else {
			fmt.Println("Não é quadrado perfeito")
		}
	}
}