## Slices

Slices wrap underlying arrays. They hold refs to the underlying array and if we assign a slice to another, both refer to the same underlying array. If a fn takes a slice as an argument, changes it makes to the slice will be visible to the caller, just like passing a pointer to the underlying array.

The capacity of a slice is the max size of the underlying array. The length of the slice can be changed as long as it fits in the capacity. Capacity can be accessed with `cap(s)` and length with `len(s)`.

We can append to a slice with `append(s, x)`.

## Two-dimensional slices

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
