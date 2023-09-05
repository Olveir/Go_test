package main

import "fmt"

func main() { //Faça um algoritmo que imprima os números de 1 a 100, substituindo os múltiplos de 3 pela palavra "Fizz" e
	// os múltiplos de 5 pela palavra "Buzz". Para os números que são múltiplos de ambos, utilize a palavra "FizzBuzz".

	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(num)
		}
	}
}
