# Módulo 01 :: Fundamentos de Go

Base del lenguaje. Sin esto el resto no fluye.

## Por qué Go

Go fue diseñado en Google (2007, liberado en 2009) por Robert Griesemer, Rob Pike y Ken Thompson para resolver tres problemas concretos: compilación lenta, concurrencia compleja y dependencias enredadas. El resultado es un lenguaje **compilado**, **fuertemente tipado**, con **garbage collector** y **concurrencia nativa** vía goroutines.

Es especialmente popular en backends de alto tráfico, APIs REST, microservicios y herramientas de infraestructura (Docker y Kubernetes están escritos en Go).

## Anatomía de un programa Go

```go
package main

import "fmt"

func main() {
    fmt.Println("hola")
}
```

Tres reglas no negociables:

1. Todo archivo `.go` empieza con una declaración `package`.
2. El paquete `main` con función `main()` es lo único ejecutable.
3. Lo que importes lo usas, lo que declares lo usas. Si no, no compila.

## Contenido del módulo

| Ejemplo | Tema | Comando |
|---------|------|---------|
| 01-hello | Hello world, paquetes, main | `go run ./01-fundamentos/01-hello` |
| 02-variables | Declaración, inferencia, zero value | `go run ./01-fundamentos/02-variables` |
| 03-constantes | const, iota, constantes tipadas | `go run ./01-fundamentos/03-constantes` |
| 04-tipos | int, float, string, bool, rune, byte | `go run ./01-fundamentos/04-tipos` |
| 05-operadores | Aritméticos, lógicos, comparación | `go run ./01-fundamentos/05-operadores` |
| 06-condicionales | if, else, switch | `go run ./01-fundamentos/06-condicionales` |
| 07-bucles | for clásico, while-style, for-range, break, continue | `go run ./01-fundamentos/07-bucles` |

## Convenciones de Go desde día uno

| Concepto | Regla |
|----------|-------|
| Nombres | `camelCase` para privados, `PascalCase` para exportados |
| Llaves | Siempre en la misma línea: `if x {`, nunca en la siguiente |
| Imports | Agrupados: estándar primero, luego terceros, luego internos |
| Formato | `gofmt` no es opcional: el formato lo decide el compilador |
| Errores | Se retornan, no se lanzan. Nada de try/catch |

## Antes de empezar

```bash
go version
```

Debe responder con algo como `go version go1.22.x`. Si no, instala desde https://go.dev/dl/.

## Ejercicios sugeridos

1. Programa que recibe tu nombre por consola y saluda.
2. Calculadora simple con switch (suma, resta, multiplica, divide).
3. FizzBuzz: imprime 1 al 100, pero múltiplos de 3 → "Fizz", de 5 → "Buzz", de ambos → "FizzBuzz".
