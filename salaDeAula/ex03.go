package main

import "fmt"

func main() {

	var salario, kw, custoKw, consumoKW float64
	fmt.Scan(&salario, &kw)
	custoKw = (salario * 0.7) / 100
	consumoKW = custoKw * kw
	fmt.Printf("custo por KW %.2f\ncusto por consumo %.2f\ncusto com desconto %.2f\n", custoKw, consumoKW, consumoKW-(consumoKW*0.1))
}
