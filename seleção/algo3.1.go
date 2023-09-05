package main

import (
	"fmt"
)

func main() { //Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

	var num int

	fmt.Println("insira o numero.")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é impar.")
	}
}
