## Formatting

### Basics

Go lets the machine handle most of the formatting.

`gofmt` will standardize the formatting of our code. I.e no need to spend time formatting structs:

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt` will line up the columns:

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

Some formatting details:

- **Indentation** - Uses tabs
- **Line length** - No line length limit
- **Parentheses** - Control sturctures don't need parentheses

### Comments

We can add comments in the following way:

- Using **/\* \*/** for multi-line comments
- Using **//** for single-line comments
