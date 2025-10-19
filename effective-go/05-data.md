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

### Arrays

Arrays are primaly used as building blocks for slices.

In Go, arrays are values: assigning an array to another copies all elements. If we pass an array to a function, the function will get a copy of the array, not a pointer to it. The size of the array is part of its type, [10]int and [20]int are different types.

In general, we should avoid using arrays and use slices instead.

### Slices

Slices wrap underlying arrays. They hold refs to the underlying array and if we assign a slice to another, both refer to the same underlying array. If a fn takes a slice as an argument, changes it makes to the slice will be visible to the caller, just like passing a pointer to the underlying array.

The capacity of a slice is the max size of the underlying array. The length of the slice can be changed as long as it fits in the capacity. Capacity can be accessed with `cap(s)` and length with `len(s)`.

We can append to a slice with `append(s, x)`.

### Two-dimensional slices

Go arrays and slices are one-dimensional. To create a 2D array or slice, we need a slice of slices or an array of arrays.

```go
type Transform [3][3]float64 // 3x3 array, an array of arrays
type LinesOfText [][]byte   // a slice of byte
```

Since slices are variable-length, each inner slice can have a different length.

```go
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
```

To allocate a 2D slice, we can allocate each slice independently, or we can allocate a single array and point the slices to it.

Each slice independently:

```go
// Allocate the top-level slice.
picture := make([][]uint8, YSize) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

Pointing to a single array:

```go
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

### Maps

Maps associate values of one type (the key) with values of another type (the value). Keys can be any type with equality operator (basically, anything but a slice).

Similar to slices, maps hold references to an underlying data structure. If we pass a map to a fn, the changes to the map will be visible to the caller.

Maps can be constructed with colon-separated key-value pairs:

```go
var timeZone = map[string]int{
    "UTC": 0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

Fetching a map is like fetching a slice:

```go
offset := timeZone["EST"]
```

We can differentiate a missing entry from a zero value with the `comma ok` idiom:

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]

// A more idiomatic way to write the above:
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}

// Or if we only care about the presence:
_, present := timeZone[tz]
```

To delete a key-value pair, we use the `delete` builtin. This is safe even if the key is missing.

### Printing

- Formatted printing in Go uses a style similar to C’s printf-family, but is richer and more general. The routines live in the fmt package.
  Go
  +1

- The main functions:

  - `fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf` — take a format string.
  - `fmt.Print`, `fmt.Println`, `fmt.Fprint` etc. — don’t take a format string; they “generate a default format for each argument.” `Println` also inserts spaces between arguments and appends a newline.

- Example equivalence:

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

### Format-strings and Default Formats

- If we want default formatting (e.g., decimals for integers) we can use the “catch-all” format `%v`. It works for any value, including arrays, slices, structs, maps.
- Example:

```go
fmt.Printf("%v\n", timeZone)
```

might output something like

```go
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

- For maps, the printed output is sorted lexicographically by key.
- Additional format flags:

  - `%+v` — when printing a struct, annotates fields with their names.
    Go

  - `%#v` — prints the value in full Go syntax (including type info) for any value.

  - `%q` — for strings or []byte, prints a quoted string literal. %#q uses back-quotes when possible.

  - `%T` — prints the type of a value.

### Customizing output via String()

- If you want to control the default format for a custom type, define a method with the signature String() string on that type. Then when you print it (with %v or the default) it will use your String() result.
- Example:

```go
type T struct { a int; b float64; c string }
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
// Then fmt.Printf("%v\n", t) prints something like: 7/-2.35/"abc\tdef"
```
