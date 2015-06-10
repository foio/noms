package types

import "io"

type Blob interface {
	Value
	// TODO: Rename just Len()
	ByteLen() uint64
	Read() io.Reader
}

func NewBlob(data []byte) Blob {
	return flatBlob{data}
}