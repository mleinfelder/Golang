package main

import "fmt"

func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro
	fmt.Println("Valor da variavel inteiro: ", inteiro)
	fmt.Println("Endereco variavel ponteiro: ", &inteiro)
	fmt.Println("Valor da variavel ponteiro", *ponteiro)

}
