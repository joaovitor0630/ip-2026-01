package main

import (
	"fmt"
)

func main() {
	var (
		n1, n2, n3 float64
	)
	var alunos = [3]float64{}

	maior := []float64{}

	for i := 0; i < 3; i++ {
		fmt.Scan(&n1)
		fmt.Scan(&n2)
		fmt.Scan(&n3)
		alunos[i] = (n1 + n2 + n3) / 3
	}

	for i := 0; i < 3; i++ {
		if alunos[i] >= 6 {
			maior = append(maior, alunos[i])
		}
	}

	fmt.Printf("todas as notas são %f\n", alunos)
	fmt.Print("\n")

	fmt.Printf("as maiores notas são : \n")
	for i := 0; i < len(maior); i++ {
		fmt.Printf("%.2f\n", maior[i])
	}

}
