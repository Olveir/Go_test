package main

import (
	"fmt"
)

func main() { //Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo, dentro ou acima do
	// peso ideal, utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

	var massa float64
	var altura float64
	var s string

	fmt.Println("Insira sua massa.")
	fmt.Scan(&massa)
	fmt.Println("Insira sua altura.")
	fmt.Scan(&altura)
	fmt.Println("Insira seu sexo. Coloque M para masculino e F para feminino.")
	fmt.Scan(&s)

	imc := (massa / (altura * altura)) * 10000
	if imc < 18 {
		fmt.Println("Abaixo do peso ideal.")
	} else if s == "M" && imc >= 18 && imc <= 26 {
		fmt.Println("Peso ideal.")
	} else if s == "M" && imc >= 26 && imc <= 30 {
		fmt.Println("Acima do peso ideal.")
	} else if s == "F" && imc >= 18 && imc < 24 {
		fmt.Println("Peso ideal.")
	} else {
		fmt.Println("Acima do peso ideal.")
	}
	fmt.Println("O seu Imc:", imc)
}
