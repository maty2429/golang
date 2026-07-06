package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// =========================================================
// context + GOROUTINES: CANCELAR TRABAJO EN CURSO CORRECTAMENTE
// =========================================================
// En Contexto/ vimos context.Context con funciones simples,
// sincrónicas. Ahora lo combinamos con goroutines: el uso REAL
// más común de context es avisarle a VARIAS goroutines en
// paralelo "paren, ya no hace falta seguir".
//
// Sin esto, corrés el riesgo de un "goroutine leak": goroutines
// que quedan corriendo para siempre (consumiendo memoria) porque
// nadie les avisó que ya no hacen falta.

// procesarLote simula el procesamiento de un producto, respetando
// la cancelación del context (igual que consultarConTimeout en
// Contexto/02, pero ahora dentro de una goroutine del pool).
func procesarLote(ctx context.Context, id int, resultados chan<- string) {
	select {
	case <-time.After(50 * time.Millisecond): // simula trabajo
		resultados <- fmt.Sprintf("producto %d procesado", id)
	case <-ctx.Done():
		resultados <- fmt.Sprintf("producto %d CANCELADO: %v", id, ctx.Err())
	}
}

func main() {
	fmt.Println("=== Cancelar varias goroutines a la vez con un mismo context ===")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
	defer cancel()

	resultados := make(chan string, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			procesarLote(ctx, id, resultados)
		}(i)
	}

	wg.Wait()
	close(resultados)

	for r := range resultados {
		fmt.Println(" -", r)
	}

	// ─────────────────────────────────────────────────────────
	// CANCELAR TODO EL TRABAJO RESTANTE ANTE UN ERROR
	// ─────────────────────────────────────────────────────────
	// Patrón real: si UNA goroutine encuentra un error grave,
	// cancela el context, y TODAS las demás (que comparten el
	// mismo ctx) se enteran y pueden frenar.

	fmt.Println("\n=== Una goroutine cancela el resto ante un error ===")

	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	resultados2 := make(chan string, 5)
	var wg2 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()

			if id == 3 {
				// Este "worker" encuentra un problema grave y cancela
				// TODO el trabajo restante, no solo el suyo.
				time.Sleep(20 * time.Millisecond)
				resultados2 <- fmt.Sprintf("producto %d: ERROR FATAL, cancelando todo", id)
				cancel2()
				return
			}

			procesarLote(ctx2, id, resultados2)
		}(i)
	}

	wg2.Wait()
	close(resultados2)

	for r := range resultados2 {
		fmt.Println(" -", r)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Un mismo ctx compartido  → cancela TODAS las goroutines que lo usan")
	fmt.Println("  WithTimeout              → límite de tiempo para todo el lote")
	fmt.Println("  WithCancel + cancel()    → una goroutine puede frenar a las demás")
	fmt.Println("  Sin esto                 → riesgo de 'goroutine leaks' (quedan colgadas)")
	fmt.Println("  Patrón real              → context.Context como primer parámetro, siempre")
}
