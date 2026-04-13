package main

import "fmt"

type os struct {
	nome  string
	teste int
}

func main() {

	var ps os

	fmt.Scan(&ps.nome)
	fmt.Print(ps.nome)

}
