package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	isTriangular := false
	for i := 1; i*i*i <= n; i++ {
		if i*(i+1)*(i+2) == n {
			isTriangular = true
			break
		}
	}
	if isTriangular {
		fmt.Println("É triangular")
	} else {
		fmt.Println("Não é triangular")
	}
}