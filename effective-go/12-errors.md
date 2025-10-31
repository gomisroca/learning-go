## Errors

Library routines must often return some sort of error indication to the caller. Go's multivalue return makes this easy.

By convention, errors have type `error`, a simple built-in interface:

```go
type error interface {
    Error() string
}
```

We can implement this interface with a richer model under the cover:

```go
// PathError records an error and the operation and path that caused it.
type PathError struct {
    Op string // "open", "unlink", etc.
    Path string // Associated file
    Err error // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```
