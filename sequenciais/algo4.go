package main

//Faça um algoritmo que leia três números reais e mostre a média ponderada
//entre eles, com pesos 2, 3 e 5, respectivamente.

import "fmt"

func main() {

	var a float64
	var b float64
	var c float64

	fmt.Println("Insira o primeiro número: ")
	fmt.Scan(&a)
	fmt.Println("Insira o segundo número: ")
	fmt.Scan(&b)
	fmt.Println("Insira o terceiro número: ")
	fmt.Scan(&c)

	mediaP := ((a * 2) + (b * 3) + (c * 4)) / 10
	fmt.Println("O resultado da média deu:", mediaP)
}
