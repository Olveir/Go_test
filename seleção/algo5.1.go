package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

	var num int

	fmt.Println("Insira um número.")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("É múltiplo de 3 e 5 ao mesmo momento.")
	} else {
		fmt.Println("Não é múltiplo de 3 e 5 ao mesmo momento.")
	}
}
