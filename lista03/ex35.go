package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("0")
		return
	}
	var bin string
	for n > 0 {
		bin = fmt.Sprintf("%d", n%2) + bin
		n /= 2
	}
	fmt.Println("Base 2:", bin)
}