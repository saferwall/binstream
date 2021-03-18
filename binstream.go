package binstream

import (
	"bytes"
	"errors"
)

// CacheSize is the current size of the cache we hold when peeking around the stream
const CacheSize = 8

// BinaryStream represents an active binary stream.
type BinaryStream struct {
	pos    int
	sType  StreamType
	binary []byte
	cache  []byte
	r      *bytes.Reader
}

// New creates a new binary stream from a byte slice (can be memory mapped file).
func New(b []byte) (*BinaryStream, error) {
	binary := make([]byte, len(b))
	n := copy(binary, b)
	if len(b) != n {
		return nil, errors.New("failed to copy input buffers")
	}
	return &BinaryStream{
		pos:    0,
		sType:  MEMORY,
		binary: binary,
		cache:  make([]byte, CacheSize),
		r:      bytes.NewReader(binary),
	}, nil
}

// Len returns the current binary stream size.
func (bs *BinaryStream) Len() int {
	return bs.r.Len()
}

// Pos returns the current position within the binary stream.
func (bs *BinaryStream) Pos() int {
	return bs.pos
}

// Read len(b) bytes from the underlying stream.
func (bs *BinaryStream) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = bs.r.Read(b)
	bs.pos += n
	return
}

// ReadAt len(b) bytes from the underlying stream at a given offset.
func (bs *BinaryStream) ReadAt(b []byte, off int64) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = bs.r.ReadAt(b, off)
	bs.pos += n
	return
}
