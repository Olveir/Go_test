package main

import "fmt"

func main() { //Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos
	// Em seguida, imprima o Slice e a soma dos valores armazenados.
	var tam int
	var soma int

	fmt.Println("Insira o tamanho do slice.")
	fmt.Scan(&tam)
	s := make([]int, tam)

	for i := 1; i < tam; i++ {
		fmt.Println("Insira os valores.")
		fmt.Scan(&s[i])
	}
	for _, val := range s {
		soma += val
	}

	fmt.Println(s)
	fmt.Println(soma)
}
