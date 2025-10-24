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
