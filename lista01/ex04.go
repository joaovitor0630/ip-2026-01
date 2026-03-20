package main

import "fmt"

func main() {
	var salarioMinimo, kwGasto float64
	fmt.Scan(&salarioMinimo, &kwGasto)

	custoPorKw := (salarioMinimo * 0.70) / 100.0
	custoConsumo := custoPorKw * kwGasto
	custoComDesconto := custoConsumo * 0.90

	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}
