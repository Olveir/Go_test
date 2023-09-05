package main

import (
	"fmt"
)

func main() { //Faça um algoritmo que leia três números inteiros e mostre o menor deles.
	var num1, num2, num3 int

	fmt.Println("Insira o primeiro número:")
	fmt.Scan(&num1)
	fmt.Println("Insira o segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Insira o terceiro número")
	fmt.Scan(&num3)

	if num1 <= num2 && num2 <= num3 {
		fmt.Println("O menor número é: ", num1)
	} else if num2 <= num1 && num1 <= num3 {
		fmt.Println("O menor número é: ", num2)
	} else {
		fmt.Println("O menor número é: ", num3)
	}

}
