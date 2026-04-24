package main
import "fmt"
func main() {
	despesas := 300.0
	ingressosIniciais := 130
	var lucroMax float64
	var precoMax float64
	var ingressosMax int
	primeiro := true

	fmt.Printf("%-10s %-15s %-10s\n", "Preço(R$)", "Ingressos", "Lucro(R$)")
	for preco := 6.0; preco >= 1.0; preco -= 0.60 {
		reducoes := (6.0 - preco) / 0.60
		ingressos := ingressosIniciais + int(reducoes+0.001)*30
		lucro := preco*float64(ingressos) - despesas

		fmt.Printf("%-10.2f %-15d %-10.2f\n", preco, ingressos, lucro)

		if primeiro || lucro > lucroMax {
			lucroMax = lucro
			precoMax = preco
			ingressosMax = ingressos
			primeiro = false
		}
	}
	fmt.Println("\nLucro Máximo Esperado:", lucroMax)
	fmt.Println("Preço do Ingresso:", precoMax)
	fmt.Println("Número de Ingressos:", ingressosMax)
}