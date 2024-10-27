package main

func main() {
	var soma, resulta = sum(1, 2)
	println(soma, resulta)
}

// A função sum agora retorna dois valores, um int e um bool
func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

// Como golang não tem try/catch, é comum retornar um valor booleano para indicar se a função foi executada com sucesso
// e um valor de erro, caso tenha ocorrido algum problema
