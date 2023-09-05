package main

import "fmt"

func main() { //Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.

	for num := 0; num <= 30; num++ {
		if num%3 == 0 {
			fmt.Println(num)
		}
	}
}
