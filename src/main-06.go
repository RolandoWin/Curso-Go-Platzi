package main

import "fmt"

func main06() {
	m := make(map[string]int)
	m["José"] = 14
	m["Pedro"] = 20
	fmt.Println(m)

	//Recorrer el map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value, ok := m["José"]
	fmt.Println(value, ok)
}
