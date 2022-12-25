package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12 //no ha sido declarado anteriormente
	var altura int = 14
	var area int

	fmt.Println("area:", base)
	fmt.Println("base:", altura)
	fmt.Println("area:", area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
