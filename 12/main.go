package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Trabalhando com structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	// Composição de structs

	gabriel := Cliente{"Gabriel", 31, true, Endereco{"Rua das Flores", 123, "São Paulo", "SP"}}

	gabriel.Endereco.Cidade = "Campinas"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", gabriel.Nome, gabriel.Idade, gabriel.Ativo, gabriel.Endereco.Cidade)
}
