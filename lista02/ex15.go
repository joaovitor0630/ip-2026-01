package main

import "fmt"

func main() {

	var (
		meses = [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
		dia   int
		mes   int
		ano   int
	)

	fmt.Scan(&dia, &mes, &ano)

	fmt.Printf("%d de %s de %d", dia, meses[mes], ano)

}
