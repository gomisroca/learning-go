## Initialization

Initialization in Go is more powerful than in C. Complex structs can be built during init, and the ordering issues among initialized objs are handled correctly.

### Constants

Constants are created at compile time and can only be numbers, chars (runes), strings or bools. Because of this, expressions that define them must be constant expressions, evaluatable by the compiler. For example: 1<<3 is a constant expression, but math.Sin(math.Pi/4) is not because it involves a function call to `math.Sin` that needs to happen at run time.

Enum constants are created using the `iota` keyword.

```go
type ByteSize float64

const (
    _           = iota              // iota = 0, ignored
    KB ByteSize = 1 << (10 * iota)  // iota = 1 → 1 << 10 = 1024
    MB                              // iota = 2 → 1 << 20 = 1048576
    GB                              // iota = 3 → 1 << 30 = 1073741824
    TB
    PB
    EB
    ZB
    YB
)
```

Since we can attach methods such as `String` to any user-defined type, arbitrary values can format themselves automatically:

```go
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

fmt.Println(ByteSize(500))         // 500.00B
fmt.Println(ByteSize(1024))        // 1.00KB
fmt.Println(ByteSize(1048576))     // 1.00MB
fmt.Println(ByteSize(5368709120))  // 5.00GB
```
