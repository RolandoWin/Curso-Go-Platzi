package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main-09() {

	a := 50
	b := &a //guarda la posicion en memoria de la variable a

	fmt.Println(b)
	fmt.Println(*b) //imprime el valor guardado en la posicion de la memoria

	*b = 100       //actualiza el valor de la variable a
	fmt.Println(a) //imprime el valor a

	//instanciamos la clase
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
