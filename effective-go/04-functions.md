## Functions

### Multiple Return Values

Go's fns can return multiple values. For example:

```go
// It will grab a number from a byte slice
// and return the number and the next position, two ints
func nextInt(b []byte,i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}

// We could use it to scan numbers in a byte slice
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

### Named Result Parameters

The return of a Go fn can be given names and used as regular vars, just like incoming params are.

```go
func nextInt(b []byte, pos int) (value int, nextPos int) {}
```

In this example, the return values are named `value` and `nextPos`.

### Defer

`defer` schedules a fn to be called immediately before the fn executing the defer returns.

```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close() // f.Close will run when the fn finishes executing

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...)
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

Defer has two advantages: we won't forget to close the file, and we can place the defer very close to the open, which is more readable.

Args to the defer fn are evaluated when the defer is executed, not when the call executes. Defers are executed in LIFO order.
