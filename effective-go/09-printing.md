## Printing

- Formatted printing in Go uses a style similar to C’s printf-family, but is richer and more general. The routines live in the fmt package.
  Go
  +1

- The main functions:

  - `fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf` — take a format string.
  - `fmt.Print`, `fmt.Println`, `fmt.Fprint` etc. — don’t take a format string; they “generate a default format for each argument.” `Println` also inserts spaces between arguments and appends a newline.

- Example equivalence:

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

### Format-strings and Default Formats

- If we want default formatting (e.g., decimals for integers) we can use the “catch-all” format `%v`. It works for any value, including arrays, slices, structs, maps.
- Example:

```go
fmt.Printf("%v\n", timeZone)
```

might output something like

```go
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

- For maps, the printed output is sorted lexicographically by key.
- Additional format flags:

  - `%+v` — when printing a struct, annotates fields with their names.
    Go

  - `%#v` — prints the value in full Go syntax (including type info) for any value.

  - `%q` — for strings or []byte, prints a quoted string literal. %#q uses back-quotes when possible.

  - `%T` — prints the type of a value.

### Customizing output via String()

- If you want to control the default format for a custom type, define a method with the signature String() string on that type. Then when you print it (with %v or the default) it will use your String() result.
- Example:

```go
type T struct { a int; b float64; c string }
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
// Then fmt.Printf("%v\n", t) prints something like: 7/-2.35/"abc\tdef"
```
