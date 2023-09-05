package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
	var num int
	fmt.Println("Insira o número.")
	fmt.Scan(&num)

	for nUm := 1; nUm <= num; nUm++ {
		if num%nUm == 0 {
			fmt.Println("Os divisores são:", nUm)
		}
	}
}
