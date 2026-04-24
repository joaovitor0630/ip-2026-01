package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("0")
		return
	}
	hexChars := "0123456789ABCDEF"
	var hex string
	for n > 0 {
		hex = string(hexChars[n%16]) + hex
		n /= 16
	}
	fmt.Println("Base 16:", hex)
}