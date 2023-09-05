package main

import "fmt"

func main() { //Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e
	// verifique se esse valor está presente no Array. Imprima o resultado da busca.

	var val int
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Insira o valor que deseja encontrar.")
	fmt.Scan(&val)

	x := false
	for _, v := range a {
		if val == v {
			x = true
		}
	}
	if x == true {
		fmt.Println("O valor desejado foi encontrado. ")
	} else {
		fmt.Println("O valor desejado não foi encontrado. ")
	}
}
