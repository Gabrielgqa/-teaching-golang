package main

import "fmt"

// Trabalhando com structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	gabriel := Cliente{"Gabriel", 31, true}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", gabriel.Nome, gabriel.Idade, gabriel.Ativo)
}
