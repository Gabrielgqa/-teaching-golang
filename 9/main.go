package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13))
}

// Função que recebe uma quantidade variável de parâmetros (função variádica)
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
