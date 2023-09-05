package main

import "fmt"

func main() { //Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
	s := []int{8, 1, 2, 5, 4, 3, 9, 7, 0, 6}

	var min, max int
	min = s[0]
	max = s[0]

	for n := range s {
		if min > n {
			min = n
		} else if max < n {
			max = n
		}
	}
	fmt.Println(min, max)
}
