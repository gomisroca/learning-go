## Names

Names in Go have semantic effect: uppercase names are visible outside its package, lowercase names are private.

### Package Names

By convention, packages are given lower case, single-word names. We shouldn't worry about collisions.

The package name is the base name of its source directory. i.e if the pkg is in `src/encoding/base64`, it is imported as `encoding/base64` but has the name `base64`.

### Getters

Go doesn't have automatic support for getters and setters. We can create them, but we shouldn't name the Get `Get`. If we have a field called `owner`, we should name the getter `Owner` and the setter `SetOwner`.

```go
owner := obj.Owner() // Get owner
if owner != user {
    obj.SetOwner(user) // Set owner
}
```

### Interface Names

One-method interfaces are named by the method name plus an -er suffix: Reader, Writer, Formatter, CloseNotifier.

### MixedCaps

By convention, Go uses MixedCaps or mixedCaps, not underscores to write multi-word names.

### Semicolons

Idiomatic Go programs have semicolons only in places such as for loop clauses, to separate the initializer, condition, and continuation elements.
