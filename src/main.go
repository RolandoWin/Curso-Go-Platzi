package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		fmt.Println("i es:", i)

		if i == 8 {
			fmt.Println("Break en 8")
			break
		}
	}
}
