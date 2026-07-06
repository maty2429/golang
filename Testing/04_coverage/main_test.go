package main

import "testing"

// =========================================================
// TESTS INCOMPLETOS A PROPÓSITO
// =========================================================
// Testeamos solo 2 de las 4 ramas de AplicarDescuento (VIP y
// default). Corré:
//
//   go test -cover ./Testing/04_coverage/
//
// y vas a ver que el coverage NO es 100%: las ramas "frecuente"
// y "empleado" quedan sin ejercitar por ningún test. Ese es
// justamente el valor de -cover: te muestra el hueco.

func TestAplicarDescuentoVIP(t *testing.T) {
	resultado := AplicarDescuento(1000, "VIP")
	esperado := 800.0

	if resultado != esperado {
		t.Errorf("AplicarDescuento(1000, VIP) = %.2f; esperado %.2f", resultado, esperado)
	}
}

func TestAplicarDescuentoSinDescuento(t *testing.T) {
	resultado := AplicarDescuento(1000, "regular")
	esperado := 1000.0

	if resultado != esperado {
		t.Errorf("AplicarDescuento(1000, regular) = %.2f; esperado %.2f", resultado, esperado)
	}
}

// ─────────────────────────────────────────────────────────
// PARA LLEGAR A 100%, DESCOMENTÁ ESTO
// ─────────────────────────────────────────────────────────
// func TestAplicarDescuentoFrecuente(t *testing.T) {
// 	resultado := AplicarDescuento(1000, "frecuente")
// 	esperado := 900.0
// 	if resultado != esperado {
// 		t.Errorf("AplicarDescuento(1000, frecuente) = %.2f; esperado %.2f", resultado, esperado)
// 	}
// }
//
// func TestAplicarDescuentoEmpleado(t *testing.T) {
// 	resultado := AplicarDescuento(1000, "empleado")
// 	esperado := 500.0
// 	if resultado != esperado {
// 		t.Errorf("AplicarDescuento(1000, empleado) = %.2f; esperado %.2f", resultado, esperado)
// 	}
// }
