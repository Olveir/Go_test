package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
	var a int
	var b int
	var c int

	fmt.Println("Digite o primeiro numero: ")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&b)
	fmt.Println("Digite o terceiro número: ")
	fmt.Scan(&c)

	soma := a + b + c
	fmt.Println(soma)
}
