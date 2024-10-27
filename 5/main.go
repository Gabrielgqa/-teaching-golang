package main

import "fmt"

func main() {
	// Array de inteiros com 3 posições
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("Posição %d: %d\n", i, v)
	}

	fmt.Println(myArray[len(myArray)-1])
}
