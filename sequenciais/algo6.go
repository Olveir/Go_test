package main

import "fmt"

func main() { //Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.

	var sal float64
	fmt.Println("Insira o salário: ")
	fmt.Scan(&sal)

	aumento := (sal * 0.15) + sal
	fmt.Println("O salário com aumento fica:", aumento)
}
