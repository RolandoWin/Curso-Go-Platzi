// Stringers Personalizar el output de structs en consola
package main

import "fmt"

type pc1 struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc1) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main_10() {
	myPC := pc1{ram: 16, brand: "ASUS", disk: 100}
	fmt.Println(myPC)
}
