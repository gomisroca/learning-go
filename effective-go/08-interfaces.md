## Interfaces and Other Types

### Interfaces

Interfaces provide a way to specify the behavior of an object. If something can do _this_, then it can be used _here_.

A type can implement multiple interfaces. For example, it could implement `sort.Interface` by providing `Len()`, `Less()`, and `Swap()` methods, while also having a custom formatter:

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

### Conversions

In the above example, `Sprint()` recreates the work `Sprint()` does for slices. We could simply convert the Sequence into a plain int[] and call Sprint.

```go
func (s Sequence) Sprint() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

It is idiomatic in Go to convert the type of an expression to access a different set of methods. For example, we can reduce the initial example to:

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

### Interface Conversions and Type Assertions

Type Switches take an interface and convert it to a concrete type.
For example, we might want the actual string if the value is a string, or the return of String() if it is implemented:

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

Type Assertions takes an interface values and extracts from it a value of a specified type. If type assertion fails, it will crash the program. We use "comma, ok" to test the assertion safely:

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

### Generality

If a type exists only to implement an interface and will never have exported methods beyond it, we don't need to export the type. We export the interface instead.

In these cases, the constructor should return an interface value, not the type.

### Interfaces and Methods

Since almost anything can have methods attached, almost anything can satisfy an interface. For example, the http pkg, which defines the `Handler` interface. Any object that implements `Handler` can serve HTTP requests.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ResponseWriter` is itself an interface with methods, including `Write()`. Therefore, `http.ResponseWriter` can be used wherever an `io.Writer`can be used.

We could implement a `Handler` that counts the number of times a page is visited:

```go
// Simple counter server
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter: %d\n", *ctr)
}
```

Then we could use it in a server:

```go
import "net/http"
// ...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

We could have an internal state that needs to be notified when a page is visited, we could use a channel:

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

Finally, maybe we want to present on /args the args used when the server was started:

```go
func ArgServer() {
    fmt.Println(os.Args)
}
```

Now we need to turn it into an HTTP server. Since we can define methods for any type except pointers and interfaces, we can write a method for a function.

Following the http pkg example, first we give ArgServer the right signature:

```go
// Argument server
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, os.Args)
}
```

Now it has the same signature as `HandlerFunc`, so it can be converted to that type and use its methods:

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

Now in "/args", the handler will have value `ArgServer` and type `HandlerFunc`. The server will invoke `ServeHTTP()` of that type, with `ArgServer` as the receiver, which in turn will invoke `ArgServer()`.
