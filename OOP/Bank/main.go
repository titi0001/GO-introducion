package main

import "fmt"

type ContaCorrente struct {
	 titular string
	 numeroAgencia int
	 numeroConta int
	 saldo float64
}

func main() {

	contaThiago := ContaCorrente{
		"Thiago",
		1234,
		123456,
		100.50,
	}

	fmt.Println(contaThiago)
	


}