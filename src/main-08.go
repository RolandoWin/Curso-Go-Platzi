package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main08() {
	var myCar pk.CarPublic
	myCar.Brand = "Lamborghini"
	myCar.Year = 2022
	fmt.Println(myCar)s

	pk.PrintMessage("Hola Platzi")
}
