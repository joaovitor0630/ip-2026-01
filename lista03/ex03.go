package main
import "fmt"
func main() {
	var salarioCarlos float64
	fmt.Scan(&salarioCarlos)
	salarioJoao := salarioCarlos / 3.0
	montanteCarlos := salarioCarlos
	montanteJoao := salarioJoao
	meses := 0
	for montanteJoao < montanteCarlos {
		montanteCarlos += montanteCarlos * 0.02
		montanteJoao += montanteJoao * 0.05
		meses++
	}
	fmt.Println("Meses:", meses)
}