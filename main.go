package main

import (
	"fmt"

	"github.com/rgalicia0729/unit-tests-go/calcular"
)

func main() {
	a := 3
	b := 5
	n := 45

	total := calcular.Sumar(a, b)
	fmt.Printf("La suma de %d y %d es: %d\n", a, b, total)

	older := calcular.GetOlder(a, b)
	fmt.Printf("El mayor entre %d y %d es: %d\n", a, b, older)

	fib := calcular.Fibonacci(n)
	fmt.Printf("El fibonacci de %d es %d\n", n, fib)
}
