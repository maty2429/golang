package main

import (
	"fmt"
	"time"
)

// =========================================================
// GOROUTINES: LO MÁS DISTINTIVO DE GO
// =========================================================
// Una goroutine es una función que corre de forma CONCURRENTE
// con el resto del programa: no espera a que termine para seguir.
// Se lanza con la palabra clave "go" antes de una llamada.
//
//   go miFuncion()   // arranca en paralelo, y el código SIGUE
//                    // a la línea siguiente sin esperarla
//
// Son MUY baratas (a diferencia de los "threads" del sistema
// operativo en otros lenguajes): un programa Go puede tener miles
// de goroutines corriendo sin problema. Esta es LA herramienta
// que hace famoso a Go para programas concurrentes (servidores
// que atienden miles de conexiones a la vez, por ejemplo).

func saludarDespacio(nombre string) {
	time.Sleep(50 * time.Millisecond) // simula trabajo que tarda
	fmt.Println("  Hola,", nombre)
}

func main() {
	fmt.Println("=== Sin goroutine: todo en orden, uno espera al otro ===")

	saludarDespacio("Ana")
	saludarDespacio("Carlos")
	fmt.Println("(esto tardó el doble: cada llamada esperó a la anterior)")

	// ─────────────────────────────────────────────────────────
	// CON goroutine: el programa NO espera
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Con 'go': se lanza y el programa sigue ===")

	go saludarDespacio("Matias") // arranca... pero main() NO espera

	fmt.Println("Esta línea se imprime CASI seguro antes que el saludo de arriba")

	// ─────────────────────────────────────────────────────────
	// EL PROBLEMA: SI main() TERMINA, EL PROGRAMA TERMINA
	// ─────────────────────────────────────────────────────────
	// Esta es LA trampa número uno con goroutines: si main() se
	// termina, TODAS las goroutines en curso se cortan de golpe,
	// hayan terminado o no. Por eso, sin algo que las "espere"
	// (tema 02: sync.WaitGroup), el saludo de Matias podría no
	// llegar a imprimirse nunca.

	fmt.Println("\n=== El problema: main() no espera a las goroutines ===")
	fmt.Println("Si el programa terminara ACÁ, el saludo de Matias podría perderse")

	// Para este ejemplo, usamos Sleep como "parche" temporal (MAL,
	// mostrado a propósito): funciona, pero es frágil y lento.
	// El tema 02 muestra la forma CORRECTA de esperar.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("(esperamos con Sleep para que el saludo alcance a imprimirse)")

	// ─────────────────────────────────────────────────────────
	// VARIAS GOROUTINES A LA VEZ
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Varias goroutines en paralelo ===")

	nombres := []string{"Sofía", "Diego", "Lucía"}
	for _, nombre := range nombres {
		go saludarDespacio(nombre)
	}

	fmt.Println("Las 3 saludos se lanzaron 'a la vez' (orden de impresión no garantizado)")
	time.Sleep(100 * time.Millisecond) // de nuevo, un parche temporal

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  go miFuncion()      → la lanza en paralelo, sin esperarla")
	fmt.Println("  Muy baratas         → miles de goroutines sin problema")
	fmt.Println("  Orden NO garantizado → cuál termina primero puede variar")
	fmt.Println("  PELIGRO             → si main() termina, se cortan todas las goroutines")
	fmt.Println("  Sleep como 'espera'  → funciona para aprender, pero es un ANTI-PATRÓN")
	fmt.Println("  La forma correcta   → sync.WaitGroup (tema 02)")
}
