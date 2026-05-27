package main

import (
	"fmt"
	"strings"
)

// =========================================================
// FUNCIONES EN GO
// =========================================================
// Las funciones son bloques de código reutilizables.
// En Go, las funciones son ciudadanos de primera clase:
// pueden asignarse a variables, pasarse como argumentos,
// y retornarse desde otras funciones.
//
// Sintaxis básica:
//   func nombreFuncion(param1 tipo1, param2 tipo2) tipoRetorno {
//       return valor
//   }

// ─────────────────────────────────────────────────────────
// FUNCIONES SIMPLES
// ─────────────────────────────────────────────────────────

// Función sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola desde una función!")
}

// Función con parámetros
func sumar(a int, b int) int {
	return a + b
}

// Cuando dos parámetros son del mismo tipo, podés abreviar
func multiplicar(a, b int) int {
	return a * b
}

// Función con múltiples retornos
func dividirConResto(a, b int) (int, int) {
	return a / b, a % b
}

// ─────────────────────────────────────────────────────────
// RETORNO NOMBRADO (named return values)
// ─────────────────────────────────────────────────────────
// Los valores de retorno tienen nombre y pueden asignarse
// dentro de la función. El "naked return" los retorna todos.
func calcularCirculo(radio float64) (area float64, perimetro float64) {
	const pi = 3.14159265
	area = pi * radio * radio
	perimetro = 2 * pi * radio
	return // naked return: retorna area y perimetro
}

// ─────────────────────────────────────────────────────────
// FUNCIONES VARIÁDICAS (número variable de argumentos)
// ─────────────────────────────────────────────────────────
// El ... indica que la función acepta 0 o más argumentos del tipo dado.
// Dentro de la función, el parámetro es un slice.
func sumarTodos(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

func unirConSeparador(sep string, palabras ...string) string {
	return strings.Join(palabras, sep)
}

// ─────────────────────────────────────────────────────────
// FUNCIONES COMO VALORES (first-class functions)
// ─────────────────────────────────────────────────────────
// Las funciones pueden asignarse a variables y pasarse como parámetros.

// Tipo de función: func(int, int) int
type OperacionInt func(int, int) int

func aplicarOperacion(a, b int, op OperacionInt) int {
	return op(a, b)
}

// ─────────────────────────────────────────────────────────
// CLOSURES (funciones que "capturan" variables externas)
// ─────────────────────────────────────────────────────────
// Un closure es una función que recuerda las variables del
// contexto donde fue creada, incluso si ese contexto ya no existe.

func crearContador(inicio int) func() int {
	cuenta := inicio // esta variable es capturada por el closure
	return func() int {
		cuenta++
		return cuenta
	}
}

func crearSumador(base int) func(int) int {
	return func(n int) int {
		return base + n // "base" es capturada del contexto exterior
	}
}

// ─────────────────────────────────────────────────────────
// RECURSIÓN
// ─────────────────────────────────────────────────────────
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1) // se llama a sí misma
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// ─────────────────────────────────────────────────────────
	// LLAMADAS BÁSICAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Funciones básicas ===")
	saludar()
	fmt.Println("sumar(3, 5) =", sumar(3, 5))
	fmt.Println("multiplicar(4, 6) =", multiplicar(4, 6))

	cociente, resto := dividirConResto(17, 5)
	fmt.Printf("17 ÷ 5 = %d resto %d\n", cociente, resto)

	area, perim := calcularCirculo(5.0)
	fmt.Printf("Círculo r=5: área=%.2f, perímetro=%.2f\n", area, perim)

	// ─────────────────────────────────────────────────────────
	// FUNCIONES VARIÁDICAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Variádicas ===")
	fmt.Println("sumarTodos():", sumarTodos())
	fmt.Println("sumarTodos(1):", sumarTodos(1))
	fmt.Println("sumarTodos(1,2,3,4,5):", sumarTodos(1, 2, 3, 4, 5))

	// Expandir un slice como argumentos con ...
	nums := []int{10, 20, 30, 40}
	fmt.Println("sumarTodos(nums...):", sumarTodos(nums...)) // expande el slice

	fmt.Println(unirConSeparador(", ", "Go", "es", "genial"))
	fmt.Println(unirConSeparador(" - ", "uno", "dos", "tres"))

	// ─────────────────────────────────────────────────────────
	// FUNCIONES ANÓNIMAS Y COMO VALORES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Funciones como valores ===")

	// Función anónima asignada a variable
	cuadrado := func(n int) int {
		return n * n
	}
	fmt.Println("cuadrado(7):", cuadrado(7))

	// Función anónima ejecutada inmediatamente (IIFE)
	resultado := func(x, y int) int {
		return x*x + y*y
	}(3, 4) // se llama inmediatamente con (3, 4)
	fmt.Println("3² + 4² =", resultado)

	// Pasar función como argumento
	fmt.Println("aplicarOperacion(10, 3, sumar):", aplicarOperacion(10, 3, sumar))
	fmt.Println("aplicarOperacion(10, 3, multiplicar):", aplicarOperacion(10, 3, multiplicar))

	// Con función anónima inline
	fmt.Println("potencia:", aplicarOperacion(2, 10, func(base, exp int) int {
		resultado := 1
		for i := 0; i < exp; i++ {
			resultado *= base
		}
		return resultado
	}))

	// ─────────────────────────────────────────────────────────
	// CLOSURES EN ACCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Closures ===")

	// Cada contador tiene su propio estado independiente
	contador1 := crearContador(0)
	contador2 := crearContador(10)

	fmt.Println("contador1:", contador1(), contador1(), contador1()) // 1 2 3
	fmt.Println("contador2:", contador2(), contador2())              // 11 12
	fmt.Println("contador1:", contador1())                           // 4 (independiente de contador2)

	// Sumadores con diferentes bases
	sumar5 := crearSumador(5)
	sumar100 := crearSumador(100)

	fmt.Println("sumar5(3):", sumar5(3))     // 8
	fmt.Println("sumar100(3):", sumar100(3)) // 103

	// Caso real: middleware de logging
	fmt.Println("\n=== Closure: middleware ===")
	loggedSumar := conLog("sumar", sumar)
	loggedSumar(10, 20)

	// ─────────────────────────────────────────────────────────
	// RECURSIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Recursión ===")
	for i := 0; i <= 7; i++ {
		fmt.Printf("factorial(%d) = %d\n", i, factorial(i))
	}

	fmt.Println("\nSecuencia de Fibonacci:")
	for i := 0; i <= 8; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
	}

	// ─────────────────────────────────────────────────────────
	// FUNCIONES DE ORDEN SUPERIOR (Higher-order functions)
	// ─────────────────────────────────────────────────────────
	// Funciones que reciben o retornan funciones.
	// Son la base de la programación funcional.

	fmt.Println("\n=== Funciones de orden superior ===")

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map: transforma cada elemento
	dobles := mapSlice(numeros, func(n int) int { return n * 2 })
	fmt.Println("Dobles:", dobles)

	// Filter: filtra elementos
	pares := filterSlice(numeros, func(n int) bool { return n%2 == 0 })
	fmt.Println("Pares:", pares)

	// Reduce: acumula un resultado
	suma := reduceSlice(numeros, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Suma total:", suma)
}

// Closure que envuelve una función con logging
func conLog(nombre string, fn func(int, int) int) func(int, int) {
	return func(a, b int) {
		resultado := fn(a, b)
		fmt.Printf("[LOG] %s(%d, %d) = %d\n", nombre, a, b, resultado)
	}
}

// Map, Filter, Reduce implementados manualmente
// (En Go real se usan slices.Map/Filter/Reduce desde Go 1.21+)
func mapSlice(s []int, fn func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func filterSlice(s []int, fn func(int) bool) []int {
	var result []int
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func reduceSlice(s []int, inicial int, fn func(int, int) int) int {
	acc := inicial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}
