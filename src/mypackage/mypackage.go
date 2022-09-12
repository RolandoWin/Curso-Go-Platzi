package mypackage

import "fmt"

// CarPublic con acceso público debido a que la primera letra es mayuscula
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage
func PrintMessage(texto string) {
	fmt.Println(texto)
}
