package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.1416
	const pi2 = 3.141592
	fmt.Println(pi, pi2)
	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Aréa cuadrado
	const baseCuadrado = 10
	var areaCuadrado int = baseCuadrado * baseCuadrado

	fmt.Println("Aréa cuadrado: ", areaCuadrado)

}
