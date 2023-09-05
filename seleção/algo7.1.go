package main

import "fmt"

func main() { //Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10% se
	//o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

	var sal float64

	fmt.Println("Insira o valor do salário.")
	fmt.Scan(&sal)

	if sal < 1000 {
		fmt.Println((sal * 0.1) + sal)
	} else {
		fmt.Println((sal * 0.05) + sal)
	}
}
