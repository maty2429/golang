package main

import "fmt"

// =========================================================
// PUNTEROS COMO ÚLTIMO RECURSO
// =========================================================
// Go tiene una filosofía clara sobre los punteros:
//
//   "Usá punteros solo cuando los necesitás."
//
// En otros lenguajes (C, C++) los punteros son omnipresentes
// porque son la única forma de evitar copias costosas.
// En Go, hay muchos casos donde pasar por valor es correcto,
// más legible, más seguro, y a veces hasta más rápido.
//
// CUÁNDO SÍ usar punteros:
//   1. Necesitás modificar el original desde una función
//   2. El struct es grande y la copia es costosa
//   3. Campos opcionales (nil = ausente)
//   4. Estructuras recursivas (linked list, árbol)
//   5. Consistencia de receptores en métodos
//
// CUÁNDO NO usar punteros (la parte que se olvida):
//   1. Structs pequeños: bool, int, string, pocas palabras
//   2. Solo lectura: no hace falta *T para leer
//   3. Interfaces: ya son referencias internamente
//   4. Slices y maps: ya son referencias internamente
//   5. Valores inmutables que no deberían cambiar
//   6. Para "optimizar" sin medir: puede ser PEOR por GC pressure

// ─────────────────────────────────────────────────────────
// CASO 1: STRUCT PEQUEÑO → pasar por valor es correcto
// ─────────────────────────────────────────────────────────

type Punto struct {
	X, Y float64 // 16 bytes — PEQUEÑO
}

// Bien: receptor valor para struct pequeño de solo lectura
func (p Punto) Distancia(otro Punto) float64 {
	dx, dy := p.X-otro.X, p.Y-otro.Y
	return dx*dx + dy*dy // sin sqrt para simplificar
}

func (p Punto) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

// Innecesario: usar *Punto solo para leer
// Esto agrega complejidad sin beneficio
func distanciaPuntero(p *Punto, otro *Punto) float64 {
	dx, dy := p.X-otro.X, p.Y-otro.Y
	return dx*dx + dy*dy
}

// ─────────────────────────────────────────────────────────
// CASO 2: SLICES Y MAPS — ya son referencias
// ─────────────────────────────────────────────────────────

// MAL: *[]int es raramente necesario
func agregarMAL(s *[]int, v int) {
	*s = append(*s, v) // sintaxis fea, generalmente innecesaria
}

// BIEN: retornar el slice actualizado (como hace append)
func agregarBIEN(s []int, v int) []int {
	return append(s, v)
}

// BIEN: para maps no necesitás puntero (ya es referencia)
func registrarMAL(m *map[string]int, k string, v int) {
	(*m)[k] = v // * innecesario, feo y confuso
}

func registrarBIEN(m map[string]int, k string, v int) {
	m[k] = v // directo, el map ya es referencia
}

// ─────────────────────────────────────────────────────────
// CASO 3: PUNTERO INNECESARIO PARA "RETORNAR MÚLTIPLES"
// ─────────────────────────────────────────────────────────

// Go tiene retorno múltiple — no necesitás punteros para esto
// En C/C++ a veces se hacía: func(int *resultado1, int *resultado2)

// MAL (estilo C): usar punteros para "retornar" múltiples valores
func dividirMAL(a, b float64, cociente *float64, resto *float64) bool {
	if b == 0 {
		return false
	}
	*cociente = float64(int(a) / int(b))
	*resto = a - *cociente*b
	return true
}

// BIEN (idiomático Go): retorno múltiple + error
func dividirBIEN(a, b float64) (float64, float64, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("división por cero")
	}
	cociente := float64(int(a) / int(b))
	resto := a - cociente*b
	return cociente, resto, nil
}

// ─────────────────────────────────────────────────────────
// CASO 4: PUNTERO REAL Y NECESARIO — struct grande que muta
// ─────────────────────────────────────────────────────────

type Pedido struct {
	ID        int
	Cliente   string
	Items     []string
	Subtotal  float64
	Descuento float64
	Impuestos float64
	Total     float64
	Estado    string
	Notas     string
	Direccion string
	Metodo    string
	Tracking  string
}

// Aquí SÍ tiene sentido *Pedido:
//   - Struct grande (evitar copia)
//   - Modificamos el original
func aplicarDescuento(p *Pedido, pct float64) {
	p.Descuento = p.Subtotal * pct
	p.Total = p.Subtotal - p.Descuento + p.Impuestos
	p.Estado = "descuento aplicado"
}

func calcularImpuestos(p *Pedido, tasa float64) {
	p.Impuestos = p.Subtotal * tasa
	p.Total = p.Subtotal - p.Descuento + p.Impuestos
}

// Solo lectura: puede recibir por valor si el pedido fuera pequeño.
// Para structs grandes, usamos puntero por eficiencia.
func resumenPedido(p *Pedido) string {
	return fmt.Sprintf("Pedido #%d | %s | $%.2f | %s",
		p.ID, p.Cliente, p.Total, p.Estado)
}

// ─────────────────────────────────────────────────────────
// CASO 5: CONFUSIÓN COMÚN — pensar que puntero = más rápido
// ─────────────────────────────────────────────────────────

// Para structs pequeños, usar puntero puede ser MÁS LENTO
// porque:
//   - Agrega una indirección (el CPU debe seguir el puntero)
//   - Puede causar cache miss (el dato está en otra parte del heap)
//   - Genera más presión en el GC (más objetos en el heap)
//
// El compilador de Go puede hacer "escape analysis":
//   - Si el dato no escapa de la función → stack (rápido)
//   - Si el dato escapa (se retorna puntero) → heap (GC lo maneja)
//
// Usar puntero innecesariamente fuerza el dato al heap.

// ─────────────────────────────────────────────────────────
// TABLA DE DECISIÓN FINAL
// ─────────────────────────────────────────────────────────

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║   PUNTEROS COMO ÚLTIMO RECURSO    ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// STRUCT PEQUEÑO: valor es más limpio
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct pequeño: valor > puntero ===")

	p1 := Punto{0, 0}
	p2 := Punto{3, 4}

	// Limpio: sin &, sin *, sin nil checks
	dist := p1.Distancia(p2)
	fmt.Printf("Distancia de %s a %s: %.1f\n", p1, p2, dist)

	// Comparar con la versión con punteros (más verbose, sin ventaja)
	dist2 := distanciaPuntero(&p1, &p2)
	fmt.Printf("Con punteros (mismo resultado, más verboso): %.1f\n", dist2)

	// ─────────────────────────────────────────────────────────
	// SLICES Y MAPS: ya son referencias
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slices y maps: ya son referencias ===")

	nums := []int{1, 2, 3}
	nums = agregarBIEN(nums, 4)
	nums = agregarBIEN(nums, 5)
	fmt.Printf("Slice con agregarBIEN: %v\n", nums)

	conteos := make(map[string]int)
	registrarBIEN(conteos, "go", 100)
	registrarBIEN(conteos, "rust", 50)
	fmt.Printf("Map con registrarBIEN: %v\n", conteos)

	// ─────────────────────────────────────────────────────────
	// RETORNO MÚLTIPLE: no necesitás punteros para esto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Retorno múltiple: idiomático Go ===")

	// Mal (estilo C con punteros de salida)
	var cociente, resto float64
	ok := dividirMAL(17, 5, &cociente, &resto)
	if ok {
		fmt.Printf("dividirMAL: cociente=%.0f, resto=%.0f\n", cociente, resto)
	}

	// Bien (retorno múltiple de Go)
	c, r, err := dividirBIEN(17, 5)
	if err == nil {
		fmt.Printf("dividirBIEN: cociente=%.0f, resto=%.0f\n", c, r)
	}

	_, _, err = dividirBIEN(10, 0)
	if err != nil {
		fmt.Printf("dividirBIEN(10,0): %v\n", err)
	}

	// ─────────────────────────────────────────────────────────
	// STRUCT GRANDE: puntero SÍ tiene sentido
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct grande: puntero justificado ===")

	pedido := &Pedido{
		ID:       42,
		Cliente:  "María López",
		Items:    []string{"Notebook", "Mouse", "Teclado"},
		Subtotal: 1600.00,
		Estado:   "pendiente",
	}

	calcularImpuestos(pedido, 0.21)
	aplicarDescuento(pedido, 0.10)
	fmt.Println(resumenPedido(pedido))

	// ─────────────────────────────────────────────────────────
	// TABLA DE DECISIÓN FINAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla de decisión: ¿puntero o valor? ===")
	fmt.Println()
	fmt.Println("  SITUACIÓN                              USAR")
	fmt.Println("  ─────────────────────────────────────  ───────")
	fmt.Println("  Función necesita modificar el dato     *T")
	fmt.Println("  Struct grande (> 64 bytes aprox)       *T")
	fmt.Println("  Campo opcional (nil = ausente)         *T")
	fmt.Println("  Estructura recursiva                   *T")
	fmt.Println("  Consistencia con otros métodos *T      *T")
	fmt.Println()
	fmt.Println("  Struct pequeño, solo lectura           T")
	fmt.Println("  Slice (ya es referencia)               []T")
	fmt.Println("  Map (ya es referencia)                 map[K]V")
	fmt.Println("  Interfaz (ya es referencia)            Interface")
	fmt.Println("  Retornar múltiples valores             (T1, T2, error)")
	fmt.Println("  Valor inmutable que no debe cambiar    T")

	fmt.Println()
	fmt.Println("=== Filosofía Go sobre punteros ===")
	fmt.Println("  Go te da punteros, pero no los impone.")
	fmt.Println("  La legibilidad y simplicidad son prioridad.")
	fmt.Println("  Antes de usar *T preguntate: ¿realmente lo necesito?")
	fmt.Println()
	fmt.Println("  Si usás un puntero, el lector del código asume que:")
	fmt.Println("    → El dato PUEDE ser nil")
	fmt.Println("    → El dato PUEDE ser modificado")
	fmt.Println("    → El dato ES grande o tiene semántica de referencia")
	fmt.Println()
	fmt.Println("  Si ninguna de esas tres cosas aplica → usá valor.")
}
