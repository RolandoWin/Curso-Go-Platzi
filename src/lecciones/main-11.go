// Lección interfaces y listas de interfaces
package main

import "fmt"

//Creamos la interface
type figuras2D interface {
	area() float64
}

//creamos la clase cuadrado
type cuadrado struct {
	lado float64
}

//Creamos la clase rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

//Area del cuadrado
func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

//Area del rectangulo
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

//Función de cálculo
func calcular(f figuras2D) {
	fmt.Println("Área:", f.area())
}

func main-11() {
	myCuadrado := cuadrado{lado: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	//Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.99, true}
	fmt.Println(myInterface...)
}
