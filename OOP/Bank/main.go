package main

import (
	"fmt"
	"main/OOP/Bank/contas"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := contas.ContaCorrente{titular: "Gustavo", saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}