package main

import "fmt"

func main() {

	// Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = y - x
	fmt.Println("Resta:", result)

	//multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//división
	result = y / x
	fmt.Println("División:", result)

	//modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//incremental
	x++
	fmt.Println("Incremental:", x)

	//decremental
	x--
	fmt.Println("Decremental:", x)
}
