package main

import "fmt"

// =========================================================
// go.mod A FONDO
// =========================================================
// go.mod es el archivo que convierte una carpeta cualquiera en un
// MÓDULO de Go: la unidad que agrupa todos tus paquetes propios y
// declara qué dependencias externas usa el proyecto.
//
// Este mismo repo tiene un go.mod en la raíz. Repasemos su contenido:
//
//   module gocito
//
//   go 1.26.4
//
// Lo desglosamos línea por línea abajo.

func main() {
	fmt.Println("=== go.mod de esta biblia ===")
	fmt.Println(`  module gocito`)
	fmt.Println(`  `)
	fmt.Println(`  go 1.26.4`)

	// ─────────────────────────────────────────────────────────
	// LÍNEA "module"
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== module gocito ===")
	fmt.Println("  Define el NOMBRE del módulo (también llamado \"module path\").")
	fmt.Println("  Es el prefijo que usan TODOS los imports internos:")
	fmt.Println(`    "gocito/Paquetes/02_organizar_en_carpetas/precios"`)
	fmt.Println("  En un proyecto que vas a publicar, el nombre del módulo suele")
	fmt.Println("  ser la URL de donde vive el código, por ejemplo:")
	fmt.Println(`    module github.com/matias/mi-api`)
	fmt.Println("  así, cualquiera puede \"go get\" tu proyecto usando esa URL.")

	// ─────────────────────────────────────────────────────────
	// LÍNEA "go"
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== go 1.26.4 ===")
	fmt.Println("  Indica qué versión mínima del LENGUAJE Go usa este módulo.")
	fmt.Println("  No es \"la versión que tenés instalada\": es la versión de las")
	fmt.Println("  REGLAS del lenguaje que aplican acá. Por ejemplo, el arreglo")
	fmt.Println("  del closure en el for (Closures/04) depende de que el go.mod")
	fmt.Println("  diga 1.22 o más. Si dijera \"go 1.18\", ese comportamiento viejo")
	fmt.Println("  volvería a aplicar, aunque tengas Go 1.26 instalado.")

	// ─────────────────────────────────────────────────────────
	// CUANDO EL PROYECTO TIENE DEPENDENCIAS EXTERNAS
	// ─────────────────────────────────────────────────────────
	// Nuestro go.mod NO tiene "require" porque solo usamos la
	// librería estándar (fmt, strings, errors...). En un proyecto
	// real que use, por ejemplo, un driver de PostgreSQL, el
	// go.mod se vería así:

	fmt.Println("\n=== Cómo se ve go.mod CON dependencias externas ===")
	fmt.Println(`  module github.com/matias/mi-api`)
	fmt.Println(`  `)
	fmt.Println(`  go 1.26.4`)
	fmt.Println(`  `)
	fmt.Println(`  require (`)
	fmt.Println(`      github.com/jackc/pgx/v5 v5.6.0`)
	fmt.Println(`      github.com/joho/godotenv v1.5.1`)
	fmt.Println(`  )`)

	// ─────────────────────────────────────────────────────────
	// LOS COMANDOS QUE MANEJAN go.mod
	// ─────────────────────────────────────────────────────────
	// (No los ejecutamos acá porque modifican archivos del proyecto,
	// pero es importante que sepas qué hace cada uno)

	fmt.Println("\n=== Comandos del día a día ===")
	fmt.Println(`  go mod init gocito   → crea un go.mod nuevo (se hace UNA vez)`)
	fmt.Println(`  go get paquete@v1.2  → agrega/actualiza una dependencia`)
	fmt.Println(`  go mod tidy          → limpia dependencias no usadas y agrega`)
	fmt.Println(`                          las que falten, según tus imports`)
	fmt.Println(`  go build ./...       → compila todo el módulo`)

	// ─────────────────────────────────────────────────────────
	// go.sum: EL ARCHIVO HERMANO DE go.mod
	// ─────────────────────────────────────────────────────────
	// Cuando hay dependencias externas, aparece también go.sum: un
	// registro con los "hashes" (huellas digitales) exactos de cada
	// versión descargada, para garantizar que SIEMPRE se instale
	// exactamente el mismo código (seguridad + reproducibilidad).
	// No se edita a mano, Go lo mantiene solo.

	fmt.Println("\n=== go.sum ===")
	fmt.Println("  Se genera junto a go.mod cuando hay dependencias externas.")
	fmt.Println("  Garantiza que todos los que clonen el repo instalen EXACTAMENTE")
	fmt.Println("  el mismo código de cada dependencia. No se edita a mano.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  go.mod        → convierte la carpeta en un módulo de Go")
	fmt.Println("  module X      → nombre/ruta base para todos los imports internos")
	fmt.Println("  go X.Y.Z      → versión mínima de las REGLAS del lenguaje")
	fmt.Println("  require       → lista de dependencias externas y su versión")
	fmt.Println("  go.sum        → huellas digitales para instalaciones reproducibles")
	fmt.Println("  go mod tidy   → mantiene go.mod/go.sum sincronizados con tu código")
}
