# 📖 Gocito — mi biblia de Go

Biblioteca personal de estudio de Go, desde cero hasta punteros.
Cada carpeta es **un tema**: un `main.go` ejecutable, con comentarios en español
que explican qué hace cada cosa, ejemplos prácticos y un resumen al final.

## Cómo ejecutar cualquier tema

Desde la raíz del proyecto:

```bash
go run ./Fundamentos/01_variables
go run ./Strings/05_concatenar_y_builder
go run ./Punteros/08_receptores_metodos
```

La idea es **leer el código de arriba a abajo** como una lección, y después
ejecutarlo para ver la salida. Jugá: cambiá valores, descomentá los ejemplos
que rompen, mirá los errores del compilador.

## Orden de estudio sugerido

1. **Fundamentos** (01 → 45): la base completa del lenguaje.
2. **Strings** (01 → 10): se puede intercalar después de Fundamentos 17 (maps),
   porque usa slices, maps y manejo de errores.
3. **Punteros** (01 → 12): el bloque final, requiere structs y métodos.

---

## 🧱 Fundamentos (45 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_variables` | Qué es una variable, declaración con `var`, reasignación |
| 02 | `02_constantes` | `const`, qué se puede y qué no, `iota` |
| 03 | `03_tipos_primitivos_1` | Enteros: `int`, `int8/16/32/64`, `uint` y sus rangos |
| 04 | `04_tipos_primitivos_2` | `float32/64`, `bool`, `string`, `rune` |
| 05 | `05_zero_values` | El valor por defecto de cada tipo |
| 06 | `06_conversiones_tipo` | Conversión explícita entre tipos numéricos: `float64(x)` |
| 07 | `07_constantes_tipadas` | Constantes tipadas vs sin tipo |
| 08 | `08_var_vs_walrus` | `var` vs `:=` — cuándo usar cada una |
| 09 | `09_operadores` | Aritméticos, comparación, lógicos, bitwise |
| 10 | `10_entrada_salida` | `fmt.Print/Println/Printf/Sprintf`, verbos de formato, lectura de consola |
| 11 | `11_errores_comunes` | Los errores típicos del principiante y cómo evitarlos |
| 12 | `12_multiples_variables` | Declaración y asignación múltiple, intercambio de valores |
| 13 | `13_comentarios` | `//`, `/* */` y convenciones |
| 14 | `14_control_flujo` | Panorama general del control de flujo |
| 15 | `15_funciones` | Introducción a funciones |
| 16 | `16_arrays_slices` | Arrays vs slices, `make`, `append`, `copy`, slicing, ordenar |
| 17 | `17_maps` | Crear, leer, escribir, borrar, chequear existencia |
| 18 | `18_structs` | Definir structs, campos, structs anidados |
| 19 | `19_condicional_if` | `if`, `else`, `if` con inicialización |
| 20 | `20_variables_sombra` | Shadowing: cuando una variable "tapa" a otra |
| 21 | `21_for_completo` | El `for` clásico de 3 partes |
| 22 | `22_for_condicion` | `for` con solo condición (el "while" de Go) |
| 23 | `23_for_infinito` | `for` infinito y cómo salir |
| 24 | `24_break_continue` | Cortar o saltear iteraciones |
| 25 | `25_for_range` | `range` sobre slices, strings, maps |
| 26 | `26_iterando_maps` | El orden aleatorio de los maps y cómo ordenarlos |
| 27 | `27_for_labels` | Labels para salir de loops anidados |
| 28 | `28_switch` | `switch` clásico, múltiples casos, `fallthrough` |
| 29 | `29_switch_blanco` | `switch` sin expresión (reemplazo de if/else if) |
| 30 | `30_primer_funcion` | Tu primera función paso a paso |
| 31 | `31_declarando_llamando` | Declarar y llamar funciones, orden en el archivo |
| 32 | `32_agregar_eliminar` | Agregar y eliminar elementos de slices |
| 33 | `33_parametros_struct` | Pasar structs como parámetros |
| 34 | `34_valor_referencia` | Semántica de valor vs referencia (antesala de Punteros) |
| 35 | `35_variadicos` | Funciones variádicas: `func(nums ...int)` |
| 36 | `36_cupones_variadic` | Ejercicio práctico con variádicas |
| 37 | `37_multiples_valores` | Funciones que retornan varios valores |
| 38 | `38_manejo_errores` | El patrón `(valor, error)`, crear y chequear errores |
| 39 | `39_retorno_nombrado` | Retornos con nombre |
| 40 | `40_defer` | `defer`: posponer ejecución hasta el final |
| 41 | `41_funciones_valores` | Funciones como valores (asignarlas a variables) |
| 42 | `42_funciones_parametros` | Funciones que reciben funciones |
| 43 | `43_metodos` | Métodos sobre tipos propios |
| 44 | `44_recursividad` | Funciones que se llaman a sí mismas |
| 45 | `45_descuento_recursivo` | Ejercicio integrador con recursión |

## 🔤 Strings (10 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_un_string` | Bytes inmutables, `len()` cuenta bytes, escapes, raw strings, `[]byte` |
| 02 | `02_utf8_y_runes` | Unicode, UTF-8, el tipo `rune`, bytes vs letras |
| 03 | `03_indexar_e_iterar` | `s[i]` da bytes, `for range` da runes, `[]rune` para acceso por letra |
| 04 | `04_inmutabilidad` | Por qué no se pueden modificar y cómo "modificarlos" bien |
| 05 | `05_concatenar_y_builder` | `+`, `Sprintf`, `Join` y `strings.Builder` para loops |
| 06 | `06_paquete_strings_buscar` | `Contains`, `HasPrefix/Suffix`, `Index`, `Count`, `EqualFold` |
| 07 | `07_paquete_strings_transformar` | `ToUpper/ToLower`, `Split`, `Fields`, `Trim*`, `Replace*` |
| 08 | `08_strconv_conversiones` | `Atoi`, `Itoa`, `ParseFloat/Bool/Int` — texto ↔ número con errores |
| 09 | `09_comparar_y_ordenar` | `==`, `<`, las trampas de mayúsculas y tildes, `sort` |
| 10 | `10_ejercicio_integrador` | Analizador de texto + parser de fichas CSV (usa todo lo anterior) |

## 📍 Punteros (12 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_valores_referencias` | Modelo de memoria, direcciones, valor vs referencia |
| 02 | `02_operadores` | `&` (dirección) y `*` (desreferenciar) |
| 03 | `03_ejercicio` | Ejercicio práctico con punteros |
| 04 | `04_nil_errores` | `nil`, panics por desreferenciar nil y cómo protegerse |
| 05 | `05_zero_vs_novalue` | Zero value vs "sin valor": punteros para campos opcionales |
| 06 | `06_mutabilidad` | Modificar el original desde una función |
| 07 | `07_punteros_structs` | Punteros a structs, el atajo `p.Campo` |
| 08 | `08_receptores_metodos` | Receptores valor vs puntero en métodos |
| 09 | `09_slices_referencia` | Por qué los slices ya "actúan" como referencias |
| 10 | `10_maps_referencia` | Ídem para maps: no necesitás `*map` |
| 11 | `11_interfaces_nil_trap` | La trampa de la interfaz que contiene un puntero nil |
| 12 | `12_punteros_ultimo_recurso` | Cuándo SÍ y cuándo NO usar punteros (tabla de decisión) |

---

## 🗺️ Próximos temas (hoja de ruta)

Lo que sigue cuando termine punteros, en orden sugerido:

1. **Interfaces** — el concepto más importante después de los punteros
2. **Paquetes y visibilidad** — organizar código en varios archivos/carpetas
3. **Closures y funciones anónimas** — profundizar lo de Fundamentos 41-42
4. **Errores avanzados** — `errors.Is/As`, wrapping, errores personalizados
5. **Generics** — funciones y tipos genéricos (`[T any]`)
6. **Testing** — `go test`, tests unitarios
7. **Concurrencia** — goroutines, channels, `sync` (lo más distintivo de Go)
