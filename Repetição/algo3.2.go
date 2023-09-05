package main

import "fmt"

func main() { //Faça um algoritmo que imprima os números ímpares de 1 a 19.

	for num := 1; num <= 20; num++ {
		if num%2 == 1 {
			fmt.Println(num)
		}
	}
}
