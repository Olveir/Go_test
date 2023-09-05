package main

import "fmt"

func main() { //Faça um algoritmo que leia vários números inteiros e mostre o maior deles. A leitura deve ser interrompida
	// quando for lido o valor zero.
	var num int
	max := 0
	for {
		fmt.Println("Insira os números.")
		fmt.Scan(&num)
		if num == 0 {
			break
		} else if num > max {
			max = num
		}
	}
	if max != 0 {
		fmt.Println("O maior número é: ", max)
	} else {
		fmt.Println("Não há número maior.")
	}
}
