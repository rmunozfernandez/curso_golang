package main

import (
	"fmt"
	"go/src/curso_golang_platzi/src/src/mypackage"
)

// Cada commit tiene una clase diferente
func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 2020

	fmt.Println(myCar)

	mypackage.PrintMessage("Hola Platzi")
}
