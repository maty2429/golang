package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// TAGS DE JSON: CONTROLAR CÓMO SE VE CADA CAMPO
// =========================================================
// Los "tags" son anotaciones que se escriben pegadas al tipo de
// cada campo del struct, entre backticks. Le dicen a encoding/json
// exactamente cómo tratar ese campo.
//
//   Campo string `json:"nombre_json"`
//
// Esto es MUY común porque las convenciones de nombres cambian
// entre lenguajes: Go usa PascalCase (Nombre, Precio) pero JSON
// (y JavaScript) suelen usar camelCase o snake_case (nombre,
// precio_final). Los tags traducen entre ambos mundos.

type ProductoAPI struct {
	Nombre        string  `json:"nombre"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	CodigoInterno string  `json:"-"`                   // "-": NUNCA aparece en el JSON
	Descuento     float64 `json:"descuento,omitempty"` // omitempty: se omite si es el zero value
}

func main() {
	fmt.Println("=== Tags básicos: renombrar campos ===")

	p := ProductoAPI{
		Nombre:        "Notebook",
		Precio:        450000,
		Stock:         3,
		CodigoInterno: "SKU-9921", // este campo NUNCA sale en el JSON
	}

	datos, _ := json.Marshal(p)
	fmt.Println(string(datos))
	// {"nombre":"Notebook","precio":450000,"stock":3}
	// (CodigoInterno no aparece, y Descuento tampoco porque es 0)

	// ─────────────────────────────────────────────────────────
	// omitempty EN ACCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== omitempty: se omite si el valor es el zero value ===")

	sinDescuento := ProductoAPI{Nombre: "Mouse", Precio: 8500, Stock: 20}
	conDescuento := ProductoAPI{Nombre: "Mouse", Precio: 8500, Stock: 20, Descuento: 0.10}

	d1, _ := json.Marshal(sinDescuento)
	d2, _ := json.Marshal(conDescuento)

	fmt.Println("Sin descuento:", string(d1)) // "descuento" NO aparece
	fmt.Println("Con descuento:", string(d2)) // "descuento" SÍ aparece

	// ─────────────────────────────────────────────────────────
	// UNMARSHAL TAMBIÉN RESPETA LOS TAGS
	// ─────────────────────────────────────────────────────────
	// La traducción funciona en los dos sentidos: si el JSON viene
	// con "nombre" (minúscula), Unmarshal lo mapea al campo Nombre
	// gracias al tag, aunque el campo de Go esté en mayúscula.

	fmt.Println("\n=== Unmarshal respeta los tags ===")

	textoJSON := `{"nombre":"Teclado","precio":12000,"stock":15}`
	var nuevo ProductoAPI
	json.Unmarshal([]byte(textoJSON), &nuevo)
	fmt.Printf("  %+v\n", nuevo)

	// ─────────────────────────────────────────────────────────
	// json:"-" : EXCLUIR UN CAMPO SIEMPRE
	// ─────────────────────────────────────────────────────────
	// Útil para datos sensibles (contraseñas, tokens, códigos
	// internos) que NUNCA deberían salir por una API.

	fmt.Println("\n=== json:\"-\" excluye el campo siempre ===")
	fmt.Println("CodigoInterno en Go:", p.CodigoInterno)
	fmt.Println("¿Aparece en el JSON?", "SKU-9921 no está en:", string(datos))

	// ─────────────────────────────────────────────────────────
	// SIN TAG: Go usa el nombre del campo tal cual
	// ─────────────────────────────────────────────────────────
	// Ya lo viste en el tema 01: si no ponés tag, la clave JSON
	// es EXACTAMENTE el nombre del campo de Go (con mayúscula).

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  `json:\"nombre\"`          → renombra la clave JSON")
	fmt.Println("  `json:\"x,omitempty\"`     → omite el campo si es su zero value")
	fmt.Println("  `json:\"-\"`               → excluye el campo SIEMPRE")
	fmt.Println("  Sin tag                    → usa el nombre del campo de Go tal cual")
	fmt.Println("  Unmarshal también los usa  → traduce en ambos sentidos")
}
