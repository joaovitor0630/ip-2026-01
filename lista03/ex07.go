package main
import "fmt"
func main() {
	var n int
	var soma, qtd, maior, menor, somaPares, qtdPares, qtdImpares int
	primeiro := true
	for {
		fmt.Scan(&n)
		if n == 30000 {
			break
		}
		soma += n
		qtd++
		if primeiro || n > maior {
			maior = n
		}
		if primeiro || n < menor {
			menor = n
		}
		primeiro = false
		if n%2 == 0 {
			somaPares += n
			qtdPares++
		} else {
			qtdImpares++
		}
	}
	if qtd > 0 {
		fmt.Println("Soma:", soma)
		fmt.Println("Quantidade:", qtd)
		fmt.Printf("Média: %.2f\n", float64(soma)/float64(qtd))
		fmt.Println("Maior:", maior)
		fmt.Println("Menor:", menor)
		if qtdPares > 0 {
			fmt.Printf("Média pares: %.2f\n", float64(somaPares)/float64(qtdPares))
		}
		fmt.Printf("Percentagem ímpares: %.2f%%\n", float64(qtdImpares)/float64(qtd)*100)
	}
}