package main

import "fmt"

func main() { //Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.

	var num int
	fmt.Println("Insira um número")
	fmt.Scan(&num)

	for nUm := 1; nUm <= 10; nUm++ {
		fmt.Println(nUm * num)
	}
}
