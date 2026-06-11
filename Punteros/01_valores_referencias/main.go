package main

import "fmt"

// =========================================================
// VALORES Y REFERENCIAS
// =========================================================
// Para entender los punteros, primero hay que entender cómo
// Go almacena los datos en memoria.
//
// Existen DOS formas de guardar/pasar datos:
//
//  POR VALOR:
//    - Se copia el dato completo.
//    - Modificar la copia NO afecta al original.
//    - Tipos: int, float, bool, string, struct, array.
//
//  POR REFERENCIA:
//    - Se pasa la DIRECCIÓN de memoria donde vive el dato.
//    - Modificar el destino de la referencia SÍ afecta al original.
//    - Tipos nativos: slice, map, channel, func, interface.
//    - Manuales: punteros (*T)
//
// ┌────────────────────────────────────────────────────────┐
// │              MEMORIA SIMPLIFICADA                       │
// │                                                         │
// │  STACK (rápida, tamaño fijo, se limpia automáticamente) │
// │  ┌───────────┐                                          │
// │  │ x = 42    │  ← variable local (valor en stack)       │
// │  │ y = 100   │                                          │
// │  │ p = 0x... │  ← puntero (guarda una DIRECCIÓN)        │
// │  └───────────┘                                          │
// │         │                                               │
// │         ▼                                               │
// │  HEAP (dinámica, tamaño variable, GC la limpia)         │
// │  ┌───────────┐                                          │
// │  │ 0xc00001: │  ← el dato al que apunta el puntero     │
// │  │  { ... }  │                                          │
// │  └───────────┘                                          │
// └────────────────────────────────────────────────────────┘

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║      VALORES Y REFERENCIAS        ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// PASO POR VALOR: se copia el dato
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Paso por VALOR (copia independiente) ===")

	a := 10
	b := a // b es una COPIA de a, viven en direcciones distintas
	b = 99 // modificar b NO afecta a a

	fmt.Printf("a = %d (sin cambios)\n", a)
	fmt.Printf("b = %d (copia modificada)\n", b)
	fmt.Printf("¿Misma dirección? a:%p  b:%p\n", &a, &b)

	// Lo mismo ocurre con structs
	type Punto struct{ X, Y int }

	p1 := Punto{1, 2}
	p2 := p1 // p2 es una COPIA completa de p1
	p2.X = 99

	fmt.Printf("\np1 = %v (sin cambios)\n", p1)
	fmt.Printf("p2 = %v (copia modificada)\n", p2)

	// ─────────────────────────────────────────────────────────
	// FUNCIONES RECIBEN COPIAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Funciones reciben copias por defecto ===")

	x := 42
	fmt.Printf("Antes de llamar:  x = %d\n", x)
	intentarModificar(x)
	fmt.Printf("Después de llamar: x = %d (sin cambio!)\n", x)

	// Para que la función modifique x, le pasamos su DIRECCIÓN
	modificarConPuntero(&x)
	fmt.Printf("Después con puntero: x = %d (modificado!)\n", x)

	// ─────────────────────────────────────────────────────────
	// TIPOS QUE YA SON REFERENCIAS (nativamente)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tipos que YA son referencias nativas ===")

	// SLICE: su header se copia, pero el array subyacente es compartido
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // slice2 comparte el mismo array interno
	slice2[0] = 999  // modifica el array compartido

	fmt.Printf("slice1 = %v (afectado! comparte array)\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)

	// MAP: siempre es una referencia
	mapa1 := map[string]int{"a": 1}
	mapa2 := mapa1 // mapa2 apunta al mismo map
	mapa2["b"] = 2 // modifica el map original

	fmt.Printf("\nmapa1 = %v (afectado! mismo map)\n", mapa1)
	fmt.Printf("mapa2 = %v\n", mapa2)

	// ─────────────────────────────────────────────────────────
	// DIRECCIÓN DE MEMORIA: ver dónde viven las variables
	// ─────────────────────────────────────────────────────────
	fmt.Printf("\n=== Direcciones de memoria con %%p ===\n")

	n1 := 100
	n2 := 200
	n3 := n1 // copia: nueva dirección

	fmt.Printf("n1 → dirección: %p | valor: %d\n", &n1, n1)
	fmt.Printf("n2 → dirección: %p | valor: %d\n", &n2, n2)
	fmt.Printf("n3 → dirección: %p | valor: %d (copia de n1, diferente dir)\n", &n3, n3)

	// ─────────────────────────────────────────────────────────
	// POR QUÉ IMPORTA ENTENDER ESTO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Implicaciones prácticas ===")
	fmt.Println("1. Pasar struct grande por valor → copia megabytes innecesariamente")
	fmt.Println("   Solución: pasar un puntero (*MiStruct)")
	fmt.Println()
	fmt.Println("2. Querer que una función modifique un int/struct")
	fmt.Println("   Solución: pasar un puntero (&miVariable)")
	fmt.Println()
	fmt.Println("3. Slice y map se 'comparten' entre funciones sin puntero")
	fmt.Println("   (Pero append al slice no se propaga sin retorno!)")
	fmt.Println()
	fmt.Println("4. Copias son SEGURAS: garantizan que nadie modifica tu dato")
	fmt.Println("   Usa copias cuando quieras inmutabilidad del argumento")

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla: tipo de paso por defecto ===")
	fmt.Println("  Tipo         │ Por defecto  │ Modificable desde fn │")
	fmt.Println("  ─────────────┼──────────────┼──────────────────────│")
	fmt.Println("  int, float   │ valor (copia)│ solo con *T          │")
	fmt.Println("  bool, string │ valor (copia)│ solo con *T          │")
	fmt.Println("  struct       │ valor (copia)│ solo con *T          │")
	fmt.Println("  array        │ valor (copia)│ solo con *T          │")
	fmt.Println("  slice        │ ref nativa   │ elementos sí, len no │")
	fmt.Println("  map          │ ref nativa   │ todo sí              │")
	fmt.Println("  *T (puntero) │ referencia   │ sí                   │")
}

func intentarModificar(n int) {
	n = 9999 // modifica la COPIA local
	fmt.Printf("  dentro de la función: n = %d (copia local)\n", n)
}

func modificarConPuntero(n *int) {
	*n = 9999 // modifica el ORIGINAL via el puntero
	fmt.Printf("  dentro de la función con puntero: *n = %d\n", *n)
}
