package main

import "fmt"

func main() { //Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

	var idade float64

	fmt.Println("Insira a sua idade: ")
	fmt.Scan(&idade)

	dias := idade * 365
	fmt.Println("A sua idade em dias é:", dias)
}
