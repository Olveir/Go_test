package main

import "fmt"

func main() { //Faça um algoritmo que imprima os números pares de 0 a 20.

	for num := 0; num <= 20; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}

}
