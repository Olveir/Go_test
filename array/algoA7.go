package main

import "fmt"

func main() { //Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
	// Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

	var num int
	s := []int{0, 0, 0, 0, 0}

	fmt.Println("Insira um número.")
	fmt.Scan(&num)

	x := false

	for _, n := range s {
		if n != num {
			x = true
		}
	}
	if x == true {
		s = append(s, num)
	}
	fmt.Println(s)
}
