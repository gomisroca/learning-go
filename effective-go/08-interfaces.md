## Interfaces and Other Types

### Interfaces

Interfaces provide a way to specify the behavior of an object. If something can do _this_, then it can be used _here_.

A type can implement multiple interfaces. For example, it could implement `sort.Interface` by providing `Len()`, `Less()`, and `Swap()` methods, while also having a custom formatter:

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

### Conversions

In the above example, `Sprint()` recreates the work `Sprint()` does for slices. We could simply convert the Sequence into a plain int[] and call Sprint.

```go
func (s Sequence) Sprint() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

It is idiomatic in Go to convert the type of an expression to access a different set of methods. For example, we can reduce the initial example to:

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

### Interface Conversions and Type Assertions

Type Switches take an interface and convert it to a concrete type.
For example, we might want the actual string if the value is a string, or the return of String() if it is implemented:

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

Type Assertions takes an interface values and extracts from it a value of a specified type. If type assertion fails, it will crash the program. We use "comma, ok" to test the assertion safely:

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

### Generality

If a type exists only to implement an interface and will never have exported methods beyond it, we don't need to export the type. We export the interface instead.

In these cases, the constructor should return an interface value, not the type.
