package main

import "fmt"

func main() { //Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.
	var dias float64
	var diaria float64
	fmt.Println("Insira o número de dias:")
	fmt.Scan(&dias)
	fmt.Println("Insira o número de diária:")
	fmt.Scan(&diaria)

	salario := dias * diaria
	fmt.Println("O salário é de: ", salario)
}
