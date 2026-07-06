package main

import (
	"fmt"
	"sync"
	"time"
)

// =========================================================
// sync.WaitGroup: ESPERAR A QUE VARIAS GOROUTINES TERMINEN
// =========================================================
// En el tema 01 usamos time.Sleep como "parche" para esperar a
// las goroutines. sync.WaitGroup es la forma CORRECTA: un
// contador que sabe exactamente cuándo todo el trabajo terminó,
// sin adivinar cuánto tiempo esperar.
//
// El patrón es SIEMPRE el mismo:
//
//   var wg sync.WaitGroup
//   wg.Add(1)        // "hay UNA goroutine más por esperar"
//   go func() {
//       defer wg.Done() // "esta goroutine ya terminó" (Fundamentos/40: defer)
//       // ... trabajo ...
//   }()
//   wg.Wait()        // bloquea hasta que el contador llegue a 0

func procesarPedido(id int, wg *sync.WaitGroup) {
	defer wg.Done() // se ejecuta SIEMPRE al salir, incluso si hay un error/panic

	time.Sleep(50 * time.Millisecond) // simula trabajo
	fmt.Printf("  Pedido #%d procesado\n", id)
}

func main() {
	fmt.Println("=== sync.WaitGroup: esperar goroutines correctamente ===")

	var wg sync.WaitGroup

	pedidos := []int{101, 102, 103, 104}

	for _, id := range pedidos {
		wg.Add(1)                  // avisamos: una goroutine más
		go procesarPedido(id, &wg) // OJO: pasamos &wg (puntero, Punteros/08)
	}

	fmt.Println("Lanzamos todos los pedidos, esperando a que terminen...")
	wg.Wait() // bloquea EXACTAMENTE hasta que las 4 llamen Done()
	fmt.Println("Todos los pedidos terminaron. Ahora sí seguimos con seguridad.")

	// ─────────────────────────────────────────────────────────
	// POR QUÉ PASAMOS *sync.WaitGroup (PUNTERO)
	// ─────────────────────────────────────────────────────────
	// Si pasáramos "wg" por VALOR (sin &), cada goroutine recibiría
	// su PROPIA COPIA del WaitGroup, y wg.Done() nunca afectaría al
	// wg original que main() está esperando con Wait(). Por eso
	// SIEMPRE se pasa por puntero (conecta con Punteros/08:
	// consistencia de receptores/parámetros con este tipo).

	// ─────────────────────────────────────────────────────────
	// wg.Add PUEDE LLAMARSE UNA SOLA VEZ, CON EL TOTAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Add() de una vez con el total ===")

	var wg2 sync.WaitGroup
	tareas := 5
	wg2.Add(tareas) // equivalente a llamar Add(1) cinco veces

	for i := 1; i <= tareas; i++ {
		go func(n int) {
			defer wg2.Done()
			fmt.Printf("  Tarea %d completada\n", n)
		}(i) // pasamos "i" como argumento: estilo explícito y claro,
		// aunque en Go 1.26 ya no haría falta para evitar el bug
		// clásico del for (ver Closures/04)
	}

	wg2.Wait()
	fmt.Println("Las 5 tareas terminaron")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  wg.Add(n)         → 'esperá n goroutines más'")
	fmt.Println("  defer wg.Done()   → dentro de la goroutine, avisa que terminó")
	fmt.Println("  wg.Wait()         → bloquea hasta que el contador llegue a 0")
	fmt.Println("  SIEMPRE *WaitGroup → pasar por puntero, nunca por valor")
	fmt.Println("  Reemplaza a Sleep  → espera EXACTA, no una adivinanza de tiempo")
}
