package main

import "fmt"

func main() { //Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as
	// ocorrências desse valor no Slice e imprima o resultado.
	s := []string{"au", "eu", "cachorro", "doido", "miau", "dó", "re", "porco"}

	var nome string
	fmt.Printf("Escolha uma palavra para retirar.%v\n", s)
	fmt.Scan(&nome)

	novoSli := make([]string, 0, len(s))
	for _, n := range s {
		if n != nome {
			novoSli = append(novoSli, n)
		}
	}
	fmt.Println(novoSli)
}
