## Embedding

Go doesn't have typical type-driven notion of subclassing. Instead, we "borrow" pieces of an implementation by embedding types within a struct or interface.

Interface embedding is very simple. For example, `io.Reader` and `io.Writer`:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`io` then has `io.ReadWriter` which is an interface that combines the two:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

A `ReadWriter` will be able to do what a `Reader` and `Writer` can do. It is the union of both embedded interfaces. Only interfaces can be embedded within interfaces.

Same idea with structs. `bufio` implements its own `bufio.Reader` and `bufio.Writer` structs using the interfaces from `io`. It also implements a reader/writer: it lists the types within the struct but doesn't give them fieldnames:

```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader // *bufio.Reader
    *Writer // *bufio.Writer
}
```

Embedding can also be a convenient way to share implementations. For example, here we have an embedded field alongside a regular field:

```go
type Job struct {
    Command string
    *log.Logger
}
```

The `Job` type now has the `Print`, `Printf`, and other methods of `*log.Logger`.

```go
job.Println("starting now...")
```
