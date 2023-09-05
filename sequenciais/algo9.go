package main

import "fmt"

func main() { //Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.
	var preco float64
	fmt.Println("Insira o preço: ")
	fmt.Scan(&preco)
	desconto := preco - (preco * 0.1)
	fmt.Println("O produto com desconto fica:", desconto)
}
