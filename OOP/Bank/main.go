package main

import (
	"fmt"
	// clientes "main/OOP/Bank/clientes"
	contas "main/OOP/Bank/contas"
)

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
