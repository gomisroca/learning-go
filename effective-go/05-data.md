## Data

### Allocation with new

Go has two allocation primitives: `new` and `make`. They do different things and apply to different types.

`new` allocates memory, but it does not initialize it, only zeroes it. It allocates zeroed storage and returns its address (pointer). The zero value of a type should be usable without further initialization. For example, sync.Mutex's zero value is an unlocked mutex, and bytes.Buffer's zero value is an empty buffer. We can use them transitively to initialize a new type:

```go
type SyncedBuffer struct {
    lock    sync.Mutex // unlocked mutex
    buffer  bytes.Buffer // empty buffer
}

// type SyncedBuffer is ready to be used right away:
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### Constructors and Composite Literals

Sometimes zero values are not good enough and we need an initializing constructor.

For example:

```go
func NewFile(fd int, name string) *File {
     if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

But that's a lot of boilerplate. We can use a _composite literal_ instead, which creates a new instance each time it is evaluated:

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    return &File{fd, name, nil, 0}
}
```

The fields must be in order and present. We can use field:value syntax too, in which case the order doesn't matter, and missing fields are set to their zero values:

```go
return &File{fd: fd, name: name}
```

### Allocation with make

`make` creates slices, maps and channels only. It returns an initialized (not zeroed) value of type T (not \*T).

```go
var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```
