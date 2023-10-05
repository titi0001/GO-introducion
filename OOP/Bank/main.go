package main

import (
	"fmt"
	contas "main/OOP/Bank/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
	fmt.Println(contaDoDenis.ObterSaldo())
}
