package main

func normalFunction(message string) {
	println(message)
}

// buena practica: al declarar mismos tipos de dato, solo declararlo en el último dato
func tripleArgument(a, b int, c string) {
	println(a, b, c)
}

// retorna una valor int
func returnValue(a int) int {
	return a * 2
}

// retorcida la devolución de dos datos
func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	println("Value:", value)

	value1, value2 := dobleReturn(2)
	println("value1 y value2", value1, value2)

	//cuando se devuelven varios datos de una función, se puede descartar usando el caracter _
	value1, _ = dobleReturn(3)
	println("value1: ", value1)
}
