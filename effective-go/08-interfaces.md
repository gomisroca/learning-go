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
