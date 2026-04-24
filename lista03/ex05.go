package main
import "fmt"
func main() {
	var idade, count50, count10_20, countTotal, countPeso40 int
	var altura, peso, sumAltura10_20 float64
	var resp int
	for {
		fmt.Scan(&idade, &altura, &peso)
		countTotal++
		if idade > 50 {
			count50++
		}
		if idade >= 10 && idade <= 20 {
			count10_20++
			sumAltura10_20 += altura
		}
		if peso < 40 {
			countPeso40++
		}
		fmt.Println("Deseja continuar? (1-Sim, Outro-Não)")
		fmt.Scan(&resp)
		if resp != 1 {
			break
		}
	}
	fmt.Println("Idade > 50:", count50)
	if count10_20 > 0 {
		fmt.Printf("Média altura 10-20: %.2f\n", sumAltura10_20/float64(count10_20))
	}
	if countTotal > 0 {
		fmt.Printf("Porcentagem peso < 40: %.2f%%\n", float64(countPeso40)/float64(countTotal)*100)
	}
}