## Control Structures

There is no do or while loop in Go. Only for, switch and if. Switch and if accept an optional init statement like that of for.

### If

```go
if x > 0 {
    return y
}
```

Braces are mandatory.

The optional init statement is often used to initialize a variable.

Here we init err, then check if it is nil:

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

### For

Similar to C, but it unifies for and while, and there's no do-while.

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

We can declare the index variable in the init statement:

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

We can use range to iterate over arrays and slices:

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

We can drop the value if we don't need it. We can drop the key if we don't need it by using the blank identifier `_`:

```go
// Drop the value
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

// Drop the key
for _, value := range oldMap {
    newMap[value] = true
}
```

### Switch

If the switch has no expression, it switches on true.

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

We can present cases separated by commas:

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

We can also use break and continue statements. Furthermore, if we label a loop, we can break out of a loop within the switch.

```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

### Type Switch

A switch can also be used to find the type of an interface var.

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```
