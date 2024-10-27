// Em Go, um ponteiro é uma variável que contém o endereço de memória de outra variável.
// Ponteiros são úteis para permitir que funções modifiquem variáveis fora de seu escopo local
// e para passar estruturas ou arrays grandes para funções de forma eficiente.
//
// Exemplo:
package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x // p é um ponteiro para x

	fmt.Println("Valor de x:", x)     // Saída: Valor de x: 10
	fmt.Println("Endereço de x:", &x) // Saída: Endereço de x: 0x... (algum endereço de memória)
	fmt.Println("Valor de p:", p)     // Saída: Valor de p: 0x... (mesmo endereço de memória que &x)
	fmt.Println("Valor em p:", *p)    // Saída: Valor em p: 10

	*p = 20                            // Mudando o valor no endereço de memória para o qual p aponta
	fmt.Println("Novo valor de x:", x) // Saída: Novo valor de x: 20
}

// Neste exemplo, `p` é um ponteiro para a variável `x`. O operador `&` é usado para obter o
// endereço de memória de `x`, e o operador `*` é usado para obter o valor armazenado no endereço
// de memória para o qual `p` aponta. Ao mudar o valor no endereço de memória para o qual `p` aponta,
// também mudamos o valor de `x`.
