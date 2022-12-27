package mypackage

import "fmt"

// struct publico
type CarPublic struct {
	Brand string
	Year  int

	// las variables que comienzan con minusculas son privadas
	// y las que empiezan con mayusculas son publicas
}

// struct privado
type carPrivate struct {
	brand string
	year  int
}

// Función publica
func PrintMessage(text string) {
	fmt.Println(text)
}
