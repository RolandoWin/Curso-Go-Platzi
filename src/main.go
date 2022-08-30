package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	fmt.Printf("\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println("counter:", counter)
		counter++
	}

	for i := 50; i > 10; i-- {
		fmt.Println("i disminuye: ", i)
	}

	//For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad, AND")
	}

	//With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto a número
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", value)
}
