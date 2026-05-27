package main

import (
	"fmt"
	"sort"
)

// =========================================================
// MAPS: REFERENCIA INTERNA
// =========================================================
// En Go, un map es siempre una REFERENCIA a una estructura
// interna del runtime. A diferencia de los slices (que tienen
// un header visible), los maps son directamente un puntero.
//
//   var m map[string]int   → m == nil (no inicializado)
//   m = make(map[string]int) → m apunta a la estructura interna
//
//   m2 := m   → m2 y m apuntan al MISMO map
//               no hay copia de datos como con slices
//
// Consecuencias:
//   - Leer de nil map: OK (retorna zero value)
//   - Escribir en nil map: PANIC
//   - Pasar un map a una función: la función ve el MISMO map
//   - No necesitás *map[K]V para modificar (ya es referencia)

// ─────────────────────────────────────────────────────────
// INVENTARIO: ejemplo de uso real
// ─────────────────────────────────────────────────────────

type Inventario struct {
	stock map[string]int
}

func NuevoInventario() *Inventario {
	return &Inventario{
		stock: make(map[string]int),
	}
}

func (inv *Inventario) Agregar(producto string, cantidad int) {
	inv.stock[producto] += cantidad
}

func (inv *Inventario) Retirar(producto string, cantidad int) error {
	actual, existe := inv.stock[producto]
	if !existe {
		return fmt.Errorf("producto '%s' no existe", producto)
	}
	if actual < cantidad {
		return fmt.Errorf("stock insuficiente: %d < %d", actual, cantidad)
	}
	inv.stock[producto] -= cantidad
	if inv.stock[producto] == 0 {
		delete(inv.stock, producto)
	}
	return nil
}

func (inv *Inventario) Stock(producto string) (int, bool) {
	cantidad, existe := inv.stock[producto]
	return cantidad, existe
}

func (inv *Inventario) Imprimir() {
	// Los maps no tienen orden determinista → ordenar para output consistente
	keys := make([]string, 0, len(inv.stock))
	for k := range inv.stock {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("    %-15s %d unidades\n", k, inv.stock[k])
	}
}

// ─────────────────────────────────────────────────────────
// CACHÉ: mapa como caché simple
// ─────────────────────────────────────────────────────────

type Cache struct {
	datos map[string]string
}

func NuevaCache() Cache {
	return Cache{datos: make(map[string]string)}
}

// La función recibe Cache por valor, pero datos es un map (referencia).
// Las modificaciones al map SÍ se ven afuera.
// (Esto puede ser sorprendente si venís de otros lenguajes)
func setEnCache(c Cache, clave, valor string) {
	c.datos[clave] = valor // modifica el map compartido
}

// ─────────────────────────────────────────────────────────
// CONJUNTO (set pattern con map)
// ─────────────────────────────────────────────────────────

type Conjunto[T comparable] struct {
	elementos map[T]struct{} // struct{} ocupa 0 bytes
}

func NuevoConjunto[T comparable]() *Conjunto[T] {
	return &Conjunto[T]{elementos: make(map[T]struct{})}
}

func (c *Conjunto[T]) Agregar(v T) {
	c.elementos[v] = struct{}{}
}

func (c *Conjunto[T]) Contiene(v T) bool {
	_, existe := c.elementos[v]
	return existe
}

func (c *Conjunto[T]) Largo() int {
	return len(c.elementos)
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║      MAPS: REFERENCIA INTERNA     ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// nil MAP vs make MAP
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== nil map vs map inicializado ===")

	var nilMap map[string]int
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)
	fmt.Printf("Leer de nil map: nilMap[\"x\"] = %d  (OK, retorna zero value)\n", nilMap["x"])

	// Escribir en nil map → panic (comentado para que no crashee)
	// nilMap["x"] = 1  → panic: assignment to entry in nil map

	inicializado := make(map[string]int)
	inicializado["contador"] = 1
	fmt.Printf("Mapa inicializado: %v\n", inicializado)

	// ─────────────────────────────────────────────────────────
	// COMPARTIR REFERENCIA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Compartir referencia ===")

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := m1 // m2 ES el mismo map (no copia)

	m2["c"] = 3     // modifica el map compartido
	delete(m2, "a") // borra de ambos

	fmt.Printf("m1 = %v  (afectado por cambios en m2)\n", m1)
	fmt.Printf("m2 = %v\n", m2)
	fmt.Println("  → m1 y m2 son el mismo map, no hay copia")

	// ─────────────────────────────────────────────────────────
	// PASAR MAP A FUNCIÓN: ya es referencia, no necesitás *map
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Map en funciones: no necesitás *map ===")

	contadores := map[string]int{}
	contarPalabras(contadores, []string{"go", "es", "go", "rápido", "go"})

	keys := []string{"go", "es", "rápido"}
	for _, k := range keys {
		fmt.Printf("  '%s': %d veces\n", k, contadores[k])
	}
	fmt.Println("  → La función modificó el map sin necesitar *map[string]int")

	// ─────────────────────────────────────────────────────────
	// CACHE: struct por valor pero map interno es referencia
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct por valor con map interno ===")

	cache := NuevaCache()
	// Pasamos cache por valor, pero el map interno es referencia
	setEnCache(cache, "usuario:1", "Ana García")
	setEnCache(cache, "usuario:2", "Carlos López")

	// Las claves se ven afuera porque el map es referencia
	fmt.Printf("cache después de setEnCache: %v\n", cache.datos)
	fmt.Println("  → El struct se copió por valor pero el map interno se comparte")
	fmt.Println("  ⚠️  Esto puede ser confuso: cuidado al pasar structs con maps por valor")

	// ─────────────────────────────────────────────────────────
	// INVENTARIO: uso real
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Inventario con map ===")

	inv := NuevoInventario()
	inv.Agregar("Notebook", 10)
	inv.Agregar("Mouse", 50)
	inv.Agregar("Teclado", 30)
	inv.Agregar("Mouse", 20) // suma al stock existente

	fmt.Println("Inventario inicial:")
	inv.Imprimir()

	if err := inv.Retirar("Mouse", 15); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Retiro de 15 Mouse: OK")
	}

	if cant, existe := inv.Stock("Mouse"); existe {
		fmt.Printf("Stock de Mouse: %d\n", cant)
	}

	if err := inv.Retirar("Monitor", 1); err != nil {
		fmt.Printf("Error esperado: %v\n", err)
	}

	// ─────────────────────────────────────────────────────────
	// CONJUNTO (set) con map
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Conjunto (set) con map[T]struct{} ===")

	visitados := NuevoConjunto[string]()
	urls := []string{"/home", "/about", "/home", "/contact", "/about", "/home"}

	for _, url := range urls {
		if !visitados.Contiene(url) {
			visitados.Agregar(url)
			fmt.Printf("  Primera visita a: %s\n", url)
		} else {
			fmt.Printf("  Ya visitado: %s\n", url)
		}
	}
	fmt.Printf("URLs únicas visitadas: %d\n", visitados.Largo())

	// ─────────────────────────────────────────────────────────
	// CLONAR UN MAP (copia real)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Clonar un map (copia independiente) ===")

	original := map[string]int{"a": 1, "b": 2, "c": 3}
	clon := clonarMap(original)

	clon["a"] = 999
	delete(clon, "b")

	fmt.Printf("original = %v  (intacto)\n", original)
	fmt.Printf("clon     = %v  (modificado independientemente)\n", clon)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen: maps como referencias ===")
	fmt.Println("  make(map[K]V)        → inicializa el map (necesario para escribir)")
	fmt.Println("  var m map[K]V        → nil, se puede leer pero NO escribir")
	fmt.Println("  m2 := m1             → misma referencia (no copia)")
	fmt.Println("  func f(m map[K]V)    → la función ve y modifica el mismo map")
	fmt.Println("  func f(m *map[K]V)   → raramente necesario (ya es referencia)")
	fmt.Println()
	fmt.Println("  Para copiar: iterar y copiar cada par K/V manualmente")
	fmt.Println("  Para set:    map[T]struct{} (struct{} ocupa 0 bytes)")
}

func contarPalabras(m map[string]int, palabras []string) {
	for _, p := range palabras {
		m[p]++ // modifica el map original (ya es referencia)
	}
}

func clonarMap(m map[string]int) map[string]int {
	copia := make(map[string]int, len(m))
	for k, v := range m {
		copia[k] = v
	}
	return copia
}
