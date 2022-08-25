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
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Aréa cuadrado: ", areaCuadrado)

	x := 10
	y := 53

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Módulo
	result = y % x
	fmt.Println("Módulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	y--
	fmt.Println("Decremental: ", y)

	//Reto
	//Calcular el aréa de un rectangulo, trapecio y circulo
	areaRectangulo := x * y
	fmt.Println("Aréa del rectangulo: ", areaRectangulo)

	areaTrapecio := ((x + y) / 2) * altura

	fmt.Println("Aréa del trapecio: ", areaTrapecio)

	var radio float64 = 10
	var areaCirculo float64 = pi * (radio * radio)
	fmt.Println("Aréa del circulo: ", areaCirculo)

}
