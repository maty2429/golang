package main

import "fmt"

func main() {
	// =========================================================
	// CONTROL DE FLUJO: if, for, switch
	// =========================================================

	// ─────────────────────────────────────────────────────────
	// IF / ELSE IF / ELSE
	// ─────────────────────────────────────────────────────────
	// Go no requiere paréntesis en la condición (pero sí llaves {}).
	nota := 75

	fmt.Println("=== if / else if / else ===")
	if nota >= 90 {
		fmt.Println("Calificación: A (Excelente)")
	} else if nota >= 80 {
		fmt.Println("Calificación: B (Muy bueno)")
	} else if nota >= 70 {
		fmt.Println("Calificación: C (Bueno)")
	} else if nota >= 60 {
		fmt.Println("Calificación: D (Aprobado)")
	} else {
		fmt.Println("Calificación: F (Desaprobado)")
	}

	// IF con inicialización (patrón muy común en Go)
	// La variable solo existe dentro del bloque if/else
	fmt.Println("\n=== if con inicialización ===")
	if edad := calcularEdad(2000); edad >= 18 {
		fmt.Printf("Tiene %d años, es mayor de edad\n", edad)
	} else {
		fmt.Printf("Tiene %d años, es menor de edad\n", edad)
	}
	// "edad" no existe aquí afuera

	// ─────────────────────────────────────────────────────────
	// FOR: el único bucle de Go (pero muy versátil)
	// ─────────────────────────────────────────────────────────
	// Go solo tiene "for", no "while" ni "do-while".
	// Pero con "for" podés hacer todo lo que harías con ellos.

	// 1. For clásico (como C/Java)
	fmt.Println("\n=== for clásico ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 2. For como while (solo condición)
	fmt.Println("\n=== for como while ===")
	contador := 1
	for contador <= 4 {
		fmt.Printf("contador = %d\n", contador)
		contador++
	}

	// 3. For infinito (como while(true))
	// Se rompe con "break"
	fmt.Println("\n=== for infinito con break ===")
	intentos := 0
	for {
		intentos++
		fmt.Printf("Intento #%d\n", intentos)
		if intentos >= 3 {
			fmt.Println("Máximo de intentos alcanzado")
			break
		}
	}

	// 4. For range: iterar colecciones (la forma más común)
	fmt.Println("\n=== for range sobre slice ===")
	frutas := []string{"manzana", "banana", "naranja", "uva"}
	for indice, fruta := range frutas {
		fmt.Printf("[%d] %s\n", indice, fruta)
	}

	// Solo el índice
	fmt.Println("\nSolo índices:")
	for i := range frutas {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Solo el valor (ignorar índice)
	fmt.Println("Solo valores:")
	for _, fruta := range frutas {
		fmt.Print(fruta, " ")
	}
	fmt.Println()

	// For range sobre string (da runes)
	fmt.Println("\n=== for range sobre string ===")
	for i, r := range "Hola" {
		fmt.Printf("índice %d: '%c'\n", i, r)
	}

	// For range sobre map
	fmt.Println("\n=== for range sobre map ===")
	capitales := map[string]string{
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
	}
	for pais, capital := range capitales {
		fmt.Printf("%s → %s\n", pais, capital)
	}

	// ─────────────────────────────────────────────────────────
	// BREAK Y CONTINUE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== break y continue ===")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // salta los pares, va al siguiente ciclo
		}
		if i > 7 {
			break // sale del loop cuando i > 7
		}
		fmt.Printf("i = %d (impar)\n", i)
	}

	// ─────────────────────────────────────────────────────────
	// LOOPS ANIDADOS CON LABELS (etiquetas)
	// ─────────────────────────────────────────────────────────
	// Los labels permiten hacer break/continue en un loop externo
	// desde dentro de un loop interno. Se usan con moderación.

	fmt.Println("\n=== loops anidados con labels ===")
externo:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("Saliendo del loop externo en i=%d, j=%d\n", i, j)
				break externo // sale del loop marcado "externo"
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH
	// ─────────────────────────────────────────────────────────
	// Switch en Go es más poderoso y limpio que en C/Java:
	// - No necesita "break" al final de cada case (es automático)
	// - Los casos pueden tener múltiples valores
	// - Puede comparar cualquier tipo (no solo int)
	// - Puede no tener expresión (switch vacío = cadena de if/else)

	fmt.Println("\n=== switch básico ===")
	dia := "martes"
	switch dia {
	case "lunes":
		fmt.Println("Inicio de semana")
	case "martes", "miércoles", "jueves": // múltiples valores en un case
		fmt.Println("Mitad de semana")
	case "viernes":
		fmt.Println("Casi fin de semana!")
	case "sábado", "domingo":
		fmt.Println("Fin de semana!")
	default:
		fmt.Println("Día no reconocido")
	}

	// Switch con tipos (type switch)
	fmt.Println("\n=== type switch ===")
	valores := []interface{}{42, "hola", 3.14, true, nil}
	for _, v := range valores {
		switch t := v.(type) {
		case int:
			fmt.Printf("%v es int, el doble es %d\n", t, t*2)
		case string:
			fmt.Printf("%v es string, longitud %d\n", t, len(t))
		case float64:
			fmt.Printf("%v es float64\n", t)
		case bool:
			fmt.Printf("%v es bool\n", t)
		case nil:
			fmt.Println("es nil")
		}
	}

	// Switch sin expresión (equivale a cadena de if/else if)
	fmt.Println("\n=== switch sin expresión ===")
	temp := 35
	switch {
	case temp < 0:
		fmt.Println("Bajo cero, ¡cuidado con el hielo!")
	case temp < 15:
		fmt.Println("Frío, abrigate")
	case temp < 25:
		fmt.Println("Temperatura agradable")
	case temp < 35:
		fmt.Println("Calor moderado")
	default:
		fmt.Println("¡Hace mucho calor!")
	}

	// fallthrough: pasar al siguiente case (raramente usado)
	fmt.Println("\n=== fallthrough ===")
	nivel := 2
	switch nivel {
	case 1:
		fmt.Println("Nivel 1: acceso básico")
		fallthrough
	case 2:
		fmt.Println("Nivel 2: acceso intermedio")
		fallthrough
	case 3:
		fmt.Println("Nivel 3: acceso avanzado")
	case 4:
		fmt.Println("Nivel 4: acceso total")
	}
	// Con nivel=2: imprime nivel 2, 3 (por fallthrough), pero no 4

	// ─────────────────────────────────────────────────────────
	// DEFER: ejecutar algo al salir de la función
	// ─────────────────────────────────────────────────────────
	// defer programa una función para ejecutarse cuando la función
	// actual TERMINE (ya sea por return normal o por panic).
	// Se usa mucho para limpiar recursos: cerrar archivos, conexiones, etc.

	fmt.Println("\n=== defer ===")
	procesarArchivo()
}

func calcularEdad(añoNacimiento int) int {
	return 2026 - añoNacimiento
}

func procesarArchivo() {
	fmt.Println("Abriendo archivo...")
	defer fmt.Println("Cerrando archivo (defer: se ejecuta al final)")
	defer fmt.Println("Liberando recursos (defer: también se ejecuta al final)")
	// Los defers se ejecutan en orden LIFO (último en entrar, primero en salir)

	fmt.Println("Procesando datos del archivo...")
	fmt.Println("Más procesamiento...")
	// Al salir de esta función, los defers se ejecutan:
	// primero "Liberando recursos", luego "Cerrando archivo"
}
