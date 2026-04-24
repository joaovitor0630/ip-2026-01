package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	var aprovados, exame, reprovados int
	var somaClasse float64
	for i := 0; i < n; i++ {
		var n1, n2 float64
		fmt.Scan(&n1, &n2)
		media := (n1 + n2) / 2
		somaClasse += media
		fmt.Printf("Média: %.2f - ", media)
		if media <= 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}
	if n > 0 {
		fmt.Println("Aprovados:", aprovados)
		fmt.Println("Exame:", exame)
		fmt.Println("Reprovados:", reprovados)
		fmt.Printf("Média da classe: %.2f\n", somaClasse/float64(n))
	}
}