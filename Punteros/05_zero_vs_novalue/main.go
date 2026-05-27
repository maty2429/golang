package main

import "fmt"

// =========================================================
// ZERO VALUE VS NO VALUE (ausencia de valor)
// =========================================================
// Este es un tema profundo y muy práctico en Go.
//
// ZERO VALUE: el valor "vacío" por defecto de cada tipo.
//   int     → 0
//   string  → ""
//   bool    → false
//   float   → 0.0
//
// PROBLEMA: a veces necesitamos distinguir entre:
//   - "el valor es 0"  (el usuario lo ingresó)
//   - "no hay valor"   (el usuario no lo ingresó)
//
// En Go, esto se resuelve con PUNTEROS (*T):
//   *int = nil    → "no hay valor" (campo ausente)
//   *int → 0      → "hay valor, y ese valor es 0"
//   *int → 42     → "hay valor, y ese valor es 42"

// ─────────────────────────────────────────────────────────
// CASO REAL: Perfil de usuario (campos opcionales)
// ─────────────────────────────────────────────────────────

// Sin punteros: no podemos distinguir "no ingresado" de "valor 0"
type PerfilV1 struct {
	Nombre    string
	Edad      int     // 0 podría ser "bebé" o "no ingresó edad"
	Descuento float64 // 0.0 podría ser "sin descuento" o "no especificó"
	VIP       bool    // false podría ser "no VIP" o "no especificado"
}

// Con punteros: *T = nil significa "no hay valor"
type PerfilV2 struct {
	Nombre    string
	Edad      *int     // nil = no ingresó la edad
	Descuento *float64 // nil = no tiene descuento especificado
	VIP       *bool    // nil = no sabemos si es VIP
}

// ─────────────────────────────────────────────────────────
// HELPERS para crear punteros a literales fácilmente
// ─────────────────────────────────────────────────────────
// En Go no podés hacer &42 directamente (solo se puede hacer
// & a una variable). Estos helpers lo simplifican.
func ptrInt(v int) *int         { return &v }
func ptrFloat(v float64) *float64 { return &v }
func ptrBool(v bool) *bool       { return &v }
func ptrString(v string) *string { return &v }

// ─────────────────────────────────────────────────────────
// PRODUCTO CON CAMPOS OPCIONALES
// ─────────────────────────────────────────────────────────
type Producto struct {
	Nombre      string
	Precio      float64
	Stock       *int     // nil = "no se sabe el stock"
	PrecioOferta *float64 // nil = "no está en oferta"
	Descripcion *string  // nil = "sin descripción"
}

func (p Producto) Info() string {
	stock := "sin info"
	if p.Stock != nil {
		stock = fmt.Sprintf("%d unidades", *p.Stock)
	}
	oferta := "precio regular"
	if p.PrecioOferta != nil {
		oferta = fmt.Sprintf("oferta: $%.2f", *p.PrecioOferta)
	}
	return fmt.Sprintf("%-15s | $%.2f | stock: %-12s | %s",
		p.Nombre, p.Precio, stock, oferta)
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║     ZERO VALUE VS NO VALUE        ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// EL PROBLEMA CON ZERO VALUE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El problema: zero value es ambiguo ===")

	u1 := PerfilV1{Nombre: "Ana", Edad: 0, Descuento: 0, VIP: false}
	u2 := PerfilV1{Nombre: "Carlos"} // no ingresó nada más

	fmt.Printf("Ana:    Edad=%d, Descuento=%.0f, VIP=%v\n", u1.Edad, u1.Descuento, u1.VIP)
	fmt.Printf("Carlos: Edad=%d, Descuento=%.0f, VIP=%v\n", u2.Edad, u2.Descuento, u2.VIP)
	fmt.Println("⚠️  ¿La edad 0 de Carlos es real o porque no ingresó nada?")
	fmt.Println("   ¡No podemos saberlo!")

	// ─────────────────────────────────────────────────────────
	// SOLUCIÓN CON PUNTEROS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Solución: *T para distinguir ausencia ===")

	// Ana ingresó todos los datos
	edad1 := 28
	desc1 := 0.15
	vip1 := true
	perfilAna := PerfilV2{
		Nombre:    "Ana",
		Edad:      &edad1,    // tiene valor: 28
		Descuento: &desc1,    // tiene valor: 15%
		VIP:       &vip1,     // tiene valor: true
	}

	// Carlos no ingresó datos opcionales
	perfilCarlos := PerfilV2{
		Nombre:    "Carlos",
		Edad:      nil,  // no ingresó edad
		Descuento: nil,  // no tiene descuento
		VIP:       nil,  // no sabemos
	}

	// Bot/sistema con edad explícita = 0 (podría ser un caso especial)
	edadBot := 0
	perfilBot := PerfilV2{
		Nombre: "BotUsuario",
		Edad:   &edadBot, // sí ingresó, y es 0
	}

	for _, p := range []PerfilV2{perfilAna, perfilCarlos, perfilBot} {
		fmt.Printf("\n%s:\n", p.Nombre)
		if p.Edad != nil {
			fmt.Printf("  Edad: %d (ingresada)\n", *p.Edad)
		} else {
			fmt.Println("  Edad: no ingresada")
		}
		if p.Descuento != nil {
			fmt.Printf("  Descuento: %.0f%%\n", *p.Descuento*100)
		} else {
			fmt.Println("  Descuento: sin definir")
		}
		if p.VIP != nil {
			fmt.Printf("  VIP: %v\n", *p.VIP)
		} else {
			fmt.Println("  VIP: sin determinar")
		}
	}

	// ─────────────────────────────────────────────────────────
	// PRODUCTO CON CAMPOS OPCIONALES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Producto con campos opcionales ===")

	stock10 := 10
	oferta999 := 999.00
	desc := "La mejor notebook del mercado"

	productos := []Producto{
		{
			Nombre:       "Notebook Pro",
			Precio:       1500.00,
			Stock:        &stock10,
			PrecioOferta: &oferta999,
			Descripcion:  &desc,
		},
		{
			Nombre: "Mouse",
			Precio: 25.99,
			// Stock, PrecioOferta, Descripcion = nil (no sabemos)
		},
		{
			Nombre: "Teclado",
			Precio: 75.50,
			Stock:  ptrInt(0), // sabemos que hay 0 (diferente a nil!)
		},
	}

	fmt.Println()
	for _, p := range productos {
		fmt.Println(" ", p.Info())
	}

	// ─────────────────────────────────────────────────────────
	// ACTUALIZACIÓN PARCIAL: ¿cuáles campos cambiar?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Actualización parcial con *T ===")
	// En una API REST, PATCH actualiza solo los campos enviados.
	// Con *T podemos saber exactamente qué campos actualizar.

	type ActualizarProducto struct {
		Nombre  *string  // nil = no cambiar
		Precio  *float64 // nil = no cambiar
		Stock   *int     // nil = no cambiar
	}

	// Actualización 1: solo precio
	update1 := ActualizarProducto{
		Precio: ptrFloat(1399.99),
	}

	// Actualización 2: precio y stock
	update2 := ActualizarProducto{
		Precio: ptrFloat(1200.00),
		Stock:  ptrInt(15),
	}

	aplicarActualizacion := func(p *Producto, u ActualizarProducto) {
		if u.Nombre != nil {
			p.Nombre = *u.Nombre
		}
		if u.Precio != nil {
			p.Precio = *u.Precio
		}
		if u.Stock != nil {
			p.Stock = u.Stock
		}
	}

	fmt.Printf("Antes:  %s\n", productos[0].Info())
	aplicarActualizacion(&productos[0], update1)
	fmt.Printf("Update1 (solo precio): %s\n", productos[0].Info())
	aplicarActualizacion(&productos[0], update2)
	fmt.Printf("Update2 (precio+stock): %s\n", productos[0].Info())

	// ─────────────────────────────────────────────────────────
	// HELPERS: ptrInt, ptrFloat, etc.
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Por qué necesitamos helpers ===")
	fmt.Println("En Go, NO podés escribir: &42  o  &\"hola\"")
	fmt.Println("Solo podés hacer & de una variable.")
	fmt.Println()
	fmt.Println("Por eso se crean helpers:")
	fmt.Println("  func ptrInt(v int) *int { return &v }")
	fmt.Println("  uso: campo = ptrInt(42)  ← limpio y conveniente")
	fmt.Println()
	fmt.Println("Desde Go 1.18 se puede generalizar con generics:")
	fmt.Println("  func ptr[T any](v T) *T { return &v }")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("int          → 0 siempre tiene un valor (puede ser ambiguo)")
	fmt.Println("*int = nil   → ausencia real de valor (no hay dato)")
	fmt.Println("*int → 0     → hay dato y es 0")
	fmt.Println()
	fmt.Println("Úsalo cuando necesitás distinguir:")
	fmt.Println("  - 'no ingresado' vs 'ingresado como vacío/cero'")
	fmt.Println("  - 'campo presente' vs 'campo ausente' (APIs PATCH)")
	fmt.Println("  - 'configuración no especificada' vs 'configuración en 0'")
}
