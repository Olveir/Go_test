package main

import "fmt"

func main() { //Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a seguinte tabela:
	//até 9 anos: mirim
	//10 a 13 anos: infantil
	//14 a 17 anos: juvenil
	//maiores de 18 anos: adulto

	var id int

	fmt.Println("Insira sua idade.")
	fmt.Scan(&id)

	if id <= 9 {
		fmt.Println("Mirim")
	} else if id >= 10 && id <= 13 {
		fmt.Println("Infantil")
	} else if id >= 14 && id <= 17 {
		fmt.Println("Juvenil")
	} else {
		fmt.Println("Adulto")
	}
}
