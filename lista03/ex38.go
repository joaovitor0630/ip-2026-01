package main
import (
	"fmt"
	"strconv"
)
func main() {
	var cpf string
	fmt.Scan(&cpf)
	if len(cpf) != 11 {
		fmt.Println("Inválido")
		return
	}
	soma1 := 0
	peso := 2
	for i := 8; i >= 0; i-- {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma1 += digito * peso
		peso++
	}
	resto1 := soma1 % 11
	d1 := 0
	if resto1 >= 2 {
		d1 = 11 - resto1
	}

	soma2 := 0
	peso = 2
	for i := 9; i >= 0; i-- {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma2 += digito * peso
		peso++
	}
	resto2 := soma2 % 11
	d2 := 0
	if resto2 >= 2 {
		d2 = 11 - resto2
	}

	digito1_informado, _ := strconv.Atoi(string(cpf[9]))
	digito2_informado, _ := strconv.Atoi(string(cpf[10]))

	if d1 == digito1_informado && d2 == digito2_informado {
		fmt.Println("CPF Válido")
	} else {
		fmt.Println("CPF Inválido")
	}
}