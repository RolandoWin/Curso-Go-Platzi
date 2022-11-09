package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string) {
	var textoReversa string
	texto = strings.ToLower(texto)
	for i := len(texto) - 1; i >= 0; i-- {
		textoReversa += string(texto[i])
	}
	if texto == textoReversa {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es un Palindromo")
	}
}

func main_05() {
	//slice := []string{"hola", "que", "hace"}

	// Primera forma:
	// for i, valor := range slice {
	// 	fmt.Println(i, valor)
	// }s

	// Segunda forma:
	// for _, valor := range slice {
	// 	fmt.Println( valor)
	// }

	// Tercera forma:
	// for i := range slice {
	// 	fmt.Println(i)
	// }

	isPalindromo("Ama")

}
