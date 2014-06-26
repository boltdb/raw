/*
Package raw provides utilities for mapping raw Go structs to byte slices.
*/
package raw

import (
	"unsafe"
)

// String represents an offset and pointer to a string in a byte slice.
type String struct {
	Offset uint16
	Length uint16
}

// Encode writes a string to a byte slice and updates the offset/length.
func (s *String) Encode(str string, value *[]byte) {
	s.Offset = uint16(len(*value))
	s.Length = uint16(len(str))
	*value = append(*value, []byte(str)...)
}

// Bytes returns a byte slice pointing to the string's contents.
func (s *String) Bytes(value []byte) []byte {
	return (*[0xFFFF]byte)(unsafe.Pointer(&value[s.Offset]))[:s.Length]
}

// String returns a Go string of the string value from an encoded byte slice.
func (s *String) String(value []byte) string {
	return string(s.Bytes(value))
}

// Time is a marker type for time.Time.
type Time int64
