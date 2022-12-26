package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}
	//puede que lo recorra aleatoriamente

	// Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	// Encontrar un valor que no está dentro de map
	value2, ok := m["Josep"]
	fmt.Println(value2, ok)

}
