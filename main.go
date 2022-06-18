package main

import "fmt"

//Função simples
func soma(a int, b int) int {
	return a + b
}

//Função retonando 2 valores
func soma_2(a, b int) (int, bool) {
	if a > b {
		return a + b, true
	}
	return a + b, false
}

func main() {
	resultado := soma(2, 6)
	resultado2, check := soma_2(2, 6)
	fmt.Println(resultado)
	fmt.Println(resultado2)
	fmt.Println(check)
}
