package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, mEx float64
	fmt.Scan(&id, &n1, &n2, &n3, &mEx)

	media := (n1 + n2*2 + n3*3 + mEx) / 7
	var conceito string

	switch {
	case media >= 9:
		conceito = "A"
	case media >= 7.5:
		conceito = "B"
	case media >= 6:
		conceito = "C"
	case media >= 4:
		conceito = "D"
	default:
		conceito = "E"
	}

	res := "REPROVADO"
	if conceito == "A" || conceito == "B" || conceito == "C" {
		res = "APROVADO"
	}
	fmt.Printf("Aluno %d, Conceito %s, %s\n", id, conceito, res)
}
