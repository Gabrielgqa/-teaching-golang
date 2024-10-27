package main

import "fmt"

func main() {
	// Map é uma estrutura de dados que armazena pares chave/valor
	// Onde a chave é um identificador único e o valor é o dado associado a essa chave
	salarios := map[string]int{"Gabriel": 1000, "João": 1200, "Maria": 800}

	fmt.Println(salarios)
	fmt.Println(salarios["Gabriel"])

	// inicializando um map vazio
	sal := make(map[string]int)
	// ou assim
	sal2 := map[string]int{}

	sal["Gabriel"] = 1000

	fmt.Println(sal)
	fmt.Println(sal2)

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	// Se você não precisa da chave, pode ignorá-la usando o caractere _ (blank identifier)
	for _, salario := range salarios {
		fmt.Println(salario)
	}
}
