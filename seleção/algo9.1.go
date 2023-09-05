package main

import "fmt"

func main() { //Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.

	var num1, num2, num3 int
	fmt.Println("Insira o primeiro número")
	fmt.Scan(&num1)
	fmt.Println("Insira o segundo número")
	fmt.Scan(&num2)
	fmt.Println("Insira o terceiro número")
	fmt.Scan(&num3)

	if num1 < num2 && num2 < num3 {
		fmt.Println(num1, num2, num3)
	} else if num1 < num3 && num3 < num2 {
		fmt.Println(num1, num3, num2)
	} else if num2 < num1 && num1 < num3 {
		fmt.Println(num2, num1, num3)
	} else if num2 < num3 && num3 < num1 {
		fmt.Println(num2, num3, num1)
	} else if num3 < num2 && num2 < num1 {
		fmt.Println(num3, num2, num1)
	} else {
		fmt.Println(num3, num1, num2)
	}
}
