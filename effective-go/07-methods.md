## Methods

### Pointers vs. Values

Methods can be defined for any named type. The receiver doesn't have to be a struct.

We can define an Append method on slices. First, we declare a named type to which we bind the method, and then make the receiver for the method a value of that type. If we use a pointer as its receiver, the method can overwrite the caller's slice, instead of having to return the updated slice.

```go
type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

We can make our fn satisfy the `io.Writer` interface by implementing the `Write` method:

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}

var b ByteSlice
fmt.Fprintf(&b, "This hour has %d days\n", 7)
```
