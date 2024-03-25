package main

import "fmt"

func main() {

	suma("Pedro", 2, 2, 5)

	ImprimirDatos("Hola", 28, true, 3.44)

	fmt.Println(Factorial(5))

}

func suma(nombre string, num ...int) {
	var total int
	for _, v := range num {
		total += v
	}
	fmt.Printf("El nombre del alumno %s la suma es: %d", nombre, total)

}

func ImprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T -  %V  ", dato, dato)
	}

}

//FUNCIONES RECURSIVAS
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
