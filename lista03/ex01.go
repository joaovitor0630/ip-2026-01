package main
import "fmt"
func main() {
	var base, exp int
	fmt.Scan(&base, &exp)
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	fmt.Println(res)
}