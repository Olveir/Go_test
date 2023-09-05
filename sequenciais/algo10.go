package main

import "fmt"

func main() { //Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.
	var peso float64
	fmt.Println("Insira o seu peso: ")
	fmt.Scan(&peso)
	conv := peso * 2.200
	fmt.Println("A conversão do seu peso de kg para lb é de: ", conv)
}
