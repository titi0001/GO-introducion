package main

import (
	"fmt"
	 c "main/OOP/Bank/contas"
)

func main() {

	contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}