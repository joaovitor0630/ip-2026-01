package main
import (
	"fmt"
	"strconv"
)
func main() {
	var n string
	fmt.Scan(&n)
	val, _ := strconv.ParseInt(n, 8, 64)
	fmt.Println("Base 10:", val)
}