package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Slice exibe os elementos de 0 a 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Slice exibe os elementos de 0 a 4
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Slice exibe os elementos que estão a partir da posição 2
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 100)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
