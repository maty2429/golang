package main

import "fmt"

// =========================================================
// PASO POR VALOR Y PASO POR REFERENCIA
// =========================================================
// Este es uno de los conceptos más importantes en Go.
// Entender la diferencia evita bugs muy difíciles de rastrear.
//
// PASO POR VALOR:
//   - Go copia el valor completo cuando llamás la función.
//   - Modificar el parámetro dentro de la función NO afecta
//     a la variable original.
//   - Tipos que se pasan por valor: int, float, bool, string,
//     structs, arrays.
//
// PASO POR REFERENCIA (con puntero):
//   - Pasás la DIRECCIÓN de memoria del valor original.
//   - Modificar el parámetro SÍ modifica el original.
//   - Usás & para obtener la dirección y * para desreferenciar.
//
// TIPOS QUE SON REFERENCIAS NATIVAS (sin necesidad de *):
//   - slices, maps, channels, funciones, interfaces.
//   - Estos tipos internamente YA contienen un puntero.

// ─────────────────────────────────────────────────────────
// EJEMPLO 1: int — paso por VALOR
// ─────────────────────────────────────────────────────────
func doblarPorValor(n int) {
	n = n * 2 // modifica la COPIA local, no el original
	fmt.Println("  dentro de doblarPorValor, n =", n)
}

// Para modificar el int, necesitamos pasar un puntero
func doblarPorReferencia(n *int) {
	*n = *n * 2 // * desreferencia: accede al valor apuntado
	fmt.Println("  dentro de doblarPorReferencia, *n =", *n)
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 2: struct — paso por VALOR vs PUNTERO
// ─────────────────────────────────────────────────────────
type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

// Por valor: recibe copia, el original no cambia
func aplicarIVAPorValor(p Producto) Producto {
	p.Precio *= 1.21 // modifica la copia
	return p         // retorna la copia modificada
}

// Por puntero: modifica el original directamente
func aplicarIVAPorPuntero(p *Producto) {
	p.Precio *= 1.21 // modifica el original
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 3: slice — referencia NATIVA
// ─────────────────────────────────────────────────────────
// Los slices internamente son {puntero al array, len, cap}.
// Al pasarlos a una función, el puntero al array se copia,
// pero APUNTA al mismo array → modificar elementos SÍ afecta al original.
// PERO: append() puede crear un nuevo array → el caller no ve el cambio.

func duplicarElementos(s []int) {
	for i := range s {
		s[i] *= 2 // modifica los elementos → SÍ afecta al original
	}
}

func agregarElemento(s []int, n int) []int {
	// append puede crear un nuevo array
	// por eso hay que RETORNAR el slice
	return append(s, n)
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 4: map — referencia NATIVA
// ─────────────────────────────────────────────────────────
// Los maps SIEMPRE son referencias. Modificar un map dentro
// de una función SIEMPRE afecta al original.
func agregarPrecio(precios map[string]float64, producto string, precio float64) {
	precios[producto] = precio // modifica el map original directamente
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  PASO POR VALOR vs REFERENCIA ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// INT: valor vs puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== int: valor vs puntero ===")

	precio := 100

	fmt.Println("Antes:", precio)
	doblarPorValor(precio)
	fmt.Println("Después de doblarPorValor:", precio, "← sin cambio!")

	doblarPorReferencia(&precio) // & = "dame la dirección de precio"
	fmt.Println("Después de doblarPorReferencia:", precio, "← modificado!")

	// ─────────────────────────────────────────────────────────
	// STRUCT: valor vs puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== struct: valor vs puntero ===")

	notebook := Producto{"Notebook", 1000.00, 5}
	mouse := Producto{"Mouse", 20.00, 30}

	fmt.Printf("notebook antes: $%.2f\n", notebook.Precio)

	// Por valor: recibimos la copia modificada
	notebookConIVA := aplicarIVAPorValor(notebook)
	fmt.Printf("notebook ORIGINAL sin cambios: $%.2f\n", notebook.Precio)
	fmt.Printf("notebookConIVA (nuevo): $%.2f\n", notebookConIVA.Precio)

	// Por puntero: el original se modifica
	fmt.Printf("\nmouse antes: $%.2f\n", mouse.Precio)
	aplicarIVAPorPuntero(&mouse)
	fmt.Printf("mouse ORIGINAL modificado: $%.2f\n", mouse.Precio)

	// ─────────────────────────────────────────────────────────
	// SLICE: referencia nativa
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== slice: referencia nativa ===")

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Antes:", numeros)

	duplicarElementos(numeros) // modifica los elementos
	fmt.Println("Después de duplicarElementos:", numeros, "← SÍ modificado!")

	// Pero append no se propaga sin retorno
	fmt.Printf("\nLen antes: %d\n", len(numeros))
	agregarElemento(numeros, 999) // el append no se ve afuera
	fmt.Printf("Len después (sin retorno): %d ← NO cambió!\n", len(numeros))

	numeros = agregarElemento(numeros, 999) // así sí funciona
	fmt.Printf("Len después (con retorno): %d ← SÍ cambió!\n", len(numeros))

	// ─────────────────────────────────────────────────────────
	// MAP: siempre referencia
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== map: siempre referencia ===")

	precios := map[string]float64{
		"Notebook": 1000.00,
	}
	fmt.Println("Antes:", precios)

	agregarPrecio(precios, "Mouse", 20.00) // modifica el map original
	agregarPrecio(precios, "Teclado", 60.00)

	fmt.Println("Después:", precios, "← SÍ modificado!")

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR CADA UNO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Guía de uso ===")
	fmt.Println("Por VALOR (copia):")
	fmt.Println("  - Tipos pequeños: int, float, bool")
	fmt.Println("  - Struct pequeño que no necesitás modificar")
	fmt.Println("  - Cuando querés garantizar inmutabilidad del argumento")
	fmt.Println()
	fmt.Println("Por REFERENCIA (puntero *):")
	fmt.Println("  - Struct grande (evita copiar megabytes de datos)")
	fmt.Println("  - Cuando la función NECESITA modificar el valor")
	fmt.Println("  - Cuando querés señalar que algo puede ser nil")
	fmt.Println()
	fmt.Println("Referencia NATIVA (sin *):")
	fmt.Println("  - slice, map, chan, func, interface → ya son referencias")
	fmt.Println("  - Para modificar ELEMENTOS de un slice: función sin *")
	fmt.Println("  - Para agregar/quitar con append: necesitás retornar el slice")

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla resumen ===")
	fmt.Println("Tipo       | Se pasa     | Modif. elementos | Modif. tamaño")
	fmt.Println("-----------|-------------|------------------|---------------")
	fmt.Println("int/float  | por valor   | necesita *       | N/A")
	fmt.Println("string     | por valor   | necesita *       | N/A (inmutable)")
	fmt.Println("struct     | por valor   | necesita *       | necesita *")
	fmt.Println("array      | por valor   | necesita *       | tamaño fijo")
	fmt.Println("slice      | ref nativa  | directo ✓        | retornar slice")
	fmt.Println("map        | ref nativa  | directo ✓        | directo ✓")
}
