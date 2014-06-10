package raw_test

import (
	"testing"
	"unsafe"

	. "github.com/boltdb/raw"
)

// Ensure that an event can be correctly encoded.
func TestString_Encode(t *testing.T) {
	// Encode to a byte slice.
	o := &O{MyString1: "foo", MyInt: 1000, MyString2: "bar"}
	v := o.Encode()

	// Map to a raw event and verify.
	r := ((*R)(unsafe.Pointer(&v[0])))
	if s := r.MyString1.String(v); s != "foo" {
		t.Fatalf("invalid string decode(1): %q", s)
	}
	if i := r.MyInt; i != 1000 {
		t.Fatalf("invalid int decode: %q", i)
	}
	if s := r.MyString2.String(v); s != "bar" {
		t.Fatalf("invalid string decode(1): %q", s)
	}
}

func BenchmarkString_Encode(b *testing.B) {
	o := &O{MyString1: "foo", MyInt: 1000, MyString2: "bar"}
	for i := 0; i < b.N; i++ {
		v := o.Encode()
		if len(v) == 0 {
			b.Fatalf("invalid string length: %d", len(v))
		}
	}
}

func BenchmarkString_Decode(b *testing.B) {
	o := &O{MyString1: "foo", MyInt: 1000, MyString2: "bar"}
	v := o.Encode()

	for i := 0; i < b.N; i++ {
		r := ((*R)(unsafe.Pointer(&v[0])))
		if len(r.MyString1.Bytes(v)) == 0 {
			b.Fatalf("invalid string length")
		}
	}
}

// O represents a test struct that will encode into R.
type O struct {
	MyString1 string
	MyInt     int
	MyString2 string
}

// Encode encodes an Event into a byte slice that can be read by a RawEvent.
func (o *O) Encode() []byte {
	var r R
	b := make([]byte, unsafe.Sizeof(r), int(unsafe.Sizeof(r))+len(o.MyString1)+len(o.MyString2))
	r.MyString1.Encode(o.MyString1, &b)
	r.MyInt = int64(o.MyInt)
	r.MyString2.Encode(o.MyString2, &b)
	copy(b, (*[unsafe.Sizeof(r)]byte)(unsafe.Pointer(&r))[:])
	return b
}

// R represents a raw struct.
type R struct {
	MyString1 String
	MyInt     int64
	MyString2 String
}
