package main

// declaração de variáveis de escopo global
const a = "Hello, World!"

// Go é fortemente tipado

var b bool    // por default é false
var c int     // por default é 0
var d string  // por default é ""
var e float64 // por default é 0.0
var g int = 5

// Quando uma variavel é declarada mas não usada, o compilador gera um erro
/*
 pode ser assim:
var (
	b bool
	c int
	d string
	e float64
	g int = 5
)
*/

func main() {
	// declaração de variáveis de escopo local
	//var f = "Hello, World!"
	b = true

	// pode ser assim a declaração de variáveis com inferência de tipo
	h := 10 // não pode ser usado fora de uma função
	println(a)
	println(h)
}
