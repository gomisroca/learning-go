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

### Panic

Sometimes an error is truly unrecoverable. In this case, we can use `panic` to stop the program and print an error message.

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

It is usually better to avoid `panic`. If possible, problems should be masked or worked around to keep the program running.
