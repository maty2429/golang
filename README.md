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

## 🗺️ Hoja de ruta completa: de acá a dominar Go

El objetivo NO es aprender todo (imposible), sino dominar lo necesario para
hacer **APIs REST bien hechas** y **programas reales**. Esta es la ruta, en
orden. Cada etapa se apoya en la anterior — no saltees.

### Etapa 1 — Completar el lenguaje 🧩

Lo que falta del lenguaje en sí. Sin esto, el código de cualquier API te va
a parecer chino.

| Tema | Qué es | Por qué importa |
|------|--------|-----------------|
| **Interfaces** | Contratos: "cualquier tipo que tenga estos métodos sirve". Implementación implícita, `fmt.Stringer`, `error`, type assertions, type switch | EL concepto más importante después de punteros. Todo Go gira alrededor de interfaces (`io.Reader`, `http.Handler`...) |
| **Closures y funciones anónimas** | Funciones que "recuerdan" las variables de donde nacieron | Los middlewares de una API son closures. Profundiza Fundamentos 41-42 |
| **Paquetes y visibilidad** | Organizar código en carpetas/paquetes propios; Mayúscula = público, minúscula = privado; `go.mod` a fondo | Para salir de archivos sueltos y armar proyectos de verdad |
| **Errores avanzados** | `errors.New`, `fmt.Errorf` con `%w` (wrapping), `errors.Is/As`, errores personalizados, `panic`/`recover` | Una API se diferencia por cómo maneja errores. Profundiza Fundamentos 38 |
| **Generics** | Funciones y tipos que aceptan cualquier tipo: `func Max[T cmp.Ordered](a, b T) T` | Solo lo básico: leerlos y usarlos. No hace falta dominarlos para APIs |

### Etapa 2 — Biblioteca estándar esencial 📚

Los paquetes que usa el 90% de los programas reales.

| Paquete | Para qué | Prioridad |
|---------|----------|-----------|
| **encoding/json** | `Marshal`/`Unmarshal`, tags `` `json:"nombre"` `` — convertir structs ↔ JSON | 🔴 Crítico: una API REST ES recibir y devolver JSON |
| **time** | Fechas, duraciones, formateo (el raro layout `2006-01-02`), timers | 🔴 Crítico: todo programa usa fechas |
| **os + archivos** | Leer/escribir archivos, `os.Args`, variables de entorno (`os.Getenv`) | 🔴 Crítico: config de la API viene de variables de entorno |
| **io / bufio** | Las interfaces `Reader`/`Writer` y lectura eficiente | 🟡 Importante: aparecen en TODAS las firmas de la stdlib |
| **context** | Cancelación y timeouts que viajan por las funciones | 🔴 Crítico: cada handler HTTP recibe un `context.Context` |
| **log/slog** | Logging estructurado (el logger moderno de Go) | 🟡 Importante: para saber qué pasa en producción |
| **regexp** | Expresiones regulares | 🟢 Útil: validaciones; con lo básico alcanza |

### Etapa 3 — Testing 🧪

Antes de concurrencia y APIs, porque vas a testear todo lo que sigue.

1. **go test** — tests unitarios, archivos `_test.go`, `t.Errorf`
2. **Table-driven tests** — EL patrón de testing de Go (un slice de casos + un loop)
3. **Subtests y coverage** — `t.Run`, `go test -cover`
4. **httptest** — testear handlers HTTP sin levantar el servidor (lo retomás en la etapa 5)

### Etapa 4 — Concurrencia 🔀

Lo más distintivo de Go. Para APIs no necesitás ser experto (el servidor HTTP
ya maneja la concurrencia por vos), pero SÍ entender qué pasa abajo.

1. **Goroutines** — `go func()`: miles de "hilos" baratos
2. **Channels** — comunicar goroutines; con y sin buffer; `select`
3. **sync.WaitGroup** — esperar a que terminen varias goroutines
4. **sync.Mutex** — proteger datos compartidos (y detectar races con `go test -race`)
5. **context + concurrencia** — cancelar trabajos en curso
6. **Patrones** — worker pools, pipeline (con uno o dos alcanza)

### Etapa 5 — APIs REST 🌐 (el objetivo)

Acá se junta TODO lo anterior. Go trae casi todo en la stdlib: con `net/http`
solo ya se hacen APIs profesionales.

1. **net/http servidor** — `http.HandleFunc`, `http.ListenAndServe`, qué es un handler
2. **Routing moderno** — el `ServeMux` de Go 1.22+: `GET /tareas/{id}` con métodos y parámetros nativos
3. **JSON in/out** — decodificar el body del request, responder JSON, status codes correctos (200, 201, 400, 404, 500)
4. **Middleware** — funciones que envuelven handlers: logging, recuperar panics, CORS, auth (son closures + interfaces ✨)
5. **http.Client** — consumir APIs de terceros (con timeouts SIEMPRE)
6. **Base de datos** — `database/sql` + driver de PostgreSQL (`pgx`); queries, `QueryRow`, transacciones; migrar esquemas
7. **Estructura de proyecto** — layout `cmd/` + `internal/`, separar handlers / servicios / repositorios
8. **Config** — variables de entorno, valores por defecto, secretos fuera del código
9. **Validación** — validar el input del usuario antes de tocar la DB
10. **Auth con JWT** — login, tokens, middleware de autenticación
11. **Graceful shutdown** — apagar el servidor sin cortar requests a la mitad (`context` + señales)
12. **Frameworks (opcional)** — `chi`, `Gin`, `Echo`: conocelos, pero la stdlib alcanza y los recruiters valoran que la domines

### Etapa 6 — Herramientas y producción 🚀

Para que tus programas salgan de tu máquina.

1. **Tooling** — `golangci-lint`, `go mod tidy`, `go build` para otros sistemas (cross-compile)
2. **Docker** — empaquetar tu API en una imagen (los binarios de Go son ideales para esto)
3. **Makefile** — automatizar build/test/run
4. **Deploy básico** — subir la API a algún servicio (Railway, Fly.io, un VPS)

### 💪 Proyectos para practicar (uno por etapa)

La teoría sin proyectos no se fija. Sugerencias en orden de dificultad:

1. **CLI de tareas** (etapas 1-2): agregar/listar/completar tareas, guardadas en un archivo JSON
2. **Conversor de monedas** (etapa 2): consume una API de cotizaciones con `http.Client`
3. **Acortador de URLs** (etapa 5): tu primera API REST, primero en memoria (map), después con DB
4. **API de tareas completa** (etapas 5-6): CRUD + PostgreSQL + auth JWT + tests + Docker — esto ya es nivel portfolio

> Regla de oro: cuando termines Punteros e Interfaces, empezá el proyecto 1
> AUNQUE sientas que te falta. Se aprende construyendo y volviendo a la
> teoría cuando algo no sale.
