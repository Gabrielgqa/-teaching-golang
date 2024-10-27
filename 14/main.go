package main

import "fmt"

// Definindo a interface Forma
type Forma interface {
	Area() float64
}

// Definindo o tipo Retangulo
type Retangulo struct {
	Largura, Altura float64
}

// Implementando o método Area para Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Definindo o tipo Circulo
type Circulo struct {
	Raio float64
}

// Implementando o método Area para Circulo
func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func imprimirArea(f Forma) {
	fmt.Println("Área:", f.Area())
}

func main() {
	r := Retangulo{Largura: 10, Altura: 5}
	c := Circulo{Raio: 7}

	imprimirArea(r)
	imprimirArea(c)
}
