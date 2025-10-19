## Maps

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
