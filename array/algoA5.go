package main

import "fmt"

func main() { //Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os
	// valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.

	matriz := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Insira o valor da matriz na posição [%d][%d]\n", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}
	fmt.Println("Sua matriz: ")

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", matriz[i][j])
		}
		fmt.Println()
	}
}
