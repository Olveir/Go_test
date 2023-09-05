package main

import (
	"fmt"
)

func main() { //Faça um algoritmo que leia dois números inteiros e mostre o maior deles.
	var num1 int
	var num2 int
	fmt.Println("Insira o primeiro valor: ")
	fmt.Scan(&num1)
	fmt.Println("Insira o segundo valor: ")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println("O maior número é:", num1)
	} else {
		fmt.Println("O maior número é:", num2)
	}
}
