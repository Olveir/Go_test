package main

import "fmt"

func main() { //Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles, se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.
	var num1, num2 int

	fmt.Println("Insira o primeiro número")
	fmt.Scan(&num1)
	fmt.Println("Insira o segundo número")
	fmt.Scan(&num2)

	mult := num1 * num2
	soma := num1 + num2

	if num1 >= 0 && num2 >= 0 {
		fmt.Println("O resultado foi:", mult)
	} else if num1 < 0 && num2 >= 0 {
		fmt.Println("O resultado foi:", soma)
	} else {
		fmt.Println("O resultado foi:", soma)
	}
}
