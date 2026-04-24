package main
import "fmt"
func main() {
	var id, peso float64
	var idMaisGordo, idMaisMagro float64
	var pesoMaisGordo, pesoMaisMagro float64
	for i := 0; i < 90; i++ {
		fmt.Scan(&id, &peso)
		if i == 0 {
			idMaisGordo, pesoMaisGordo = id, peso
			idMaisMagro, pesoMaisMagro = id, peso
		} else {
			if peso > pesoMaisGordo {
				pesoMaisGordo = peso
				idMaisGordo = id
			}
			if peso < pesoMaisMagro {
				pesoMaisMagro = peso
				idMaisMagro = id
			}
		}
	}
	fmt.Printf("Mais gordo - ID: %.0f, Peso: %.2f\n", idMaisGordo, pesoMaisGordo)
	fmt.Printf("Mais magro - ID: %.0f, Peso: %.2f\n", idMaisMagro, pesoMaisMagro)
}