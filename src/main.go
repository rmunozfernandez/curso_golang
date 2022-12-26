package main

import "fmt"

type car struct {
	brand string
	year  int
}

type persona struct {
	name string
	year int
}

// Cada commit tiene una clase diferente
func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	human := persona{name: "Rodrigo", year: 38}
	fmt.Println(human)

	var human2 persona
	human2.name = "Paula"
	human2.year = 40
	fmt.Println(human2)
}
