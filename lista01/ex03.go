package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)

	if n1 > 9 || n2 > 9 || n3 > 9 || n1 < 0 || n2 < 0 || n3 < 0 {
		fmt.Println("DIGITO INVALIDO")
		return
	}

	numero := n1*100 + n2*10 + n3
	quadrado := numero * numero

	fmt.Printf("%d, %d\n", numero, quadrado)
}
