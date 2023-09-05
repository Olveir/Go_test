package main

import "fmt"

func main() { //Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em
	// todas as posições do Array. Imprima o Array resultante.

	a := [6]float64{1.5, 2.5, 3.0, 4.2, 5.3, 6.1}

	var num float64

	fmt.Println("Insira um número.")
	fmt.Scan(&num)

	for n := range a {
		a[n] += num
	}
	fmt.Println(a)
}
