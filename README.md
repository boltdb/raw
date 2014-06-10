raw [![Build Status](https://drone.io/github.com/boltdb/raw/status.png)](https://drone.io/github.com/boltdb/raw/latest) [![Coverage Status](https://img.shields.io/coveralls/boltdb/raw.svg)](https://coveralls.io/r/boltdb/raw?branch=master) [![GoDoc](https://godoc.org/github.com/boltdb/raw?status.png)](https://godoc.org/github.com/boltdb/raw) ![Project status](http://img.shields.io/status/experimental.png?color=red)
===

This is simple library for working with raw Go struct data. Most of the time
it's good to serialize your data to a common encoding format (e.g. JSON,
MessagePack, Protocol Buffers) when saving and retrieving data to disk or
sending over a network. These encodings can provide a common interface for data
and support versioning and other useful features.

However, serialization comes at a cost. Converting between types and copying
memory all has overhead so when you need to go really fast, sometimes you need
to skip serialization all together.


### Usage

#### Basics

Go provides the ability to perform type conversion on byte slices to convert
them into pointers of Go types. You can do this using the `unsafe` package.
As the name suggests, it's not safe. You need to know what you're doing.

```go
// Create a byte slice with 4 bytes.
b := make([]byte, 4)

// Create an 32-bit int pointer to the first byte of the slice and set a value.
x := (*int32)(unsafe.Pointer(&b[0]))
*x = 1000

// Verify that the underlying byte slice changed.
fmt.Printf("%x\n", b)
```

This will print out the value: `e8030000` which is the hex representation of `1000`.


#### Using Raw

The primitive integer and float types in Go map directly to byte slices. However,
the string type does not. Its internal representation is publicly accessible
or guaranteed not to change between Go versions. So to map variable length
Go strings to byte slices in our code we can use the `raw.String` type:

```go
var s String
b := make([]byte, unsafe.Sizeof(s))
s.Encode("foo", &b)
copy(b, (*[unsafe.Sizeof(s)]byte)(unsafe.Pointer(&s))[:])
```

That will encode the string offset and length followed by the bytes, `"foo"`.
Then when you want to use the string, type convert the byte slice to your
`raw.String` and extract the data:

```go
s := ((*raw.String)(unsafe.Pointer(&b[0])))
fmt.Print(s.String())

// Prints: foo
```

If this seems like a lot of work just to encode a string then you'd be correct.
However, it's fast and when multiple strings are combined in a struct it allows
us to only deserialize the fields we need.


### Performance

To get an idea of the performance of this approach, please see the benchmarks
inside the test suite.

On my Intel Core i7 2.9GHz Macbook Pro, I see the following stats:

```sh
$ go test -bench=. -benchmem
PASS
BenchmarkString_Encode	10000000	       214 ns/op	      64 B/op	       3 allocs/op
BenchmarkString_Decode	500000000	         3.77 ns/op	       0 B/op	       0 allocs/op
ok  	github.com/boltdb/raw	4.635s
```

YMMV.
