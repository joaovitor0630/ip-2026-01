package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var publico float64
		var pPop, pGer, pArq, pCad float64
		fmt.Scan(&publico, &pPop, &pGer, &pArq, &pCad)

		renda := (publico * (pPop / 100) * 1.0) +
			(publico * (pGer / 100) * 5.0) +
			(publico * (pArq / 100) * 10.0) +
			(publico * (pCad / 100) * 20.0)

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
