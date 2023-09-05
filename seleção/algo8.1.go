package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).

	var num int
	fmt.Println("Insira um numero entre 1 e 7 para mostrar o dia da semana correspondente.")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Println("O dia da semana é Domingo")
	} else if num == 2 {
		fmt.Println("O dia da semana é Segunda")
	} else if num == 3 {
		fmt.Println("O dia da semana é Terça")
	} else if num == 4 {
		fmt.Println("O dia da semana é Quarta")
	} else if num == 5 {
		fmt.Println("O dia da semana é Quinta")
	} else if num == 6 {
		fmt.Println("O dia da semana é Sexta")
	} else {
		fmt.Println("O dia da semana é Sábado")
	}
}
