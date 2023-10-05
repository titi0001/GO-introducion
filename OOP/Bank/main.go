package main

import (
	"fmt"
	// clientes "main/OOP/Bank/clientes"
	contas "main/OOP/Bank/contas"
)

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)

	fmt.Println(contaDoDenis.ObterSaldo())
}
