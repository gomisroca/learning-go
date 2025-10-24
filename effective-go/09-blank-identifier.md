## The Blank Identifier

The blank identifier can be declared with any value of any type, with the value discarded harmlessly. It represents an irrelevant value.

### In Multiple Assignment

Blank identifiers in for range loops are a special case of multiple assignment.

If an assignment requires multiple values on the left side, but one of them is irrelevant, we can use the blank identifier to discard the value.

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
```

### Unused Imports and Variables

We shouldn't import a pkg or declare a variable if we aren't going to use it. However, under certain circumstances, we might want unused imports or variables. We can use the blank identifier for that:

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

By convention, these blank declarations come right after the imports and are commented.

### Import for Side Effect

Blank assignments identify code as WIP. However, sometimes we need to import a package for its side effects. We can use the blank identifier to do that:

```go
import _ "net/http/pprof"
```

This way we declare the pkg is imported for its side effects, but there is no other use for it.

### Interface Checks

If it's necessary only to ask whether a type implements an interface, without actually using the interface itself, perhaps as part of an error check, use the blank identifier to ignore the type-asserted value:

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```
