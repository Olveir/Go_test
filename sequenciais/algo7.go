package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.
	var num int
	fmt.Println("Insira o número:")
	fmt.Scan(&num)
	ant := num - 1
	suc := num + 1
	fmt.Println("O antecessor:", ant, "\nO sucessor: ", suc)
}
