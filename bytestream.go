package binstream

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// CacheSize is the current size of the cache we hold when peeking around the stream
const CacheSize = 8

// ByteStream represents an in-memory binary stream.
type ByteStream struct {
	pos    int
	sType  StreamType
	binary []byte
	cache  []byte
	r      *bytes.Reader
}

// Assert interface implementation checks
var _ Stream = (*ByteStream)(nil)

// NewByteStream creates a new binary stream from a byte slice (can be memory mapped file).
func NewByteStream(b []byte) (*ByteStream, error) {
	binary := make([]byte, len(b))
	n := copy(binary, b)
	if len(b) != n {
		return nil, errors.New("failed to copy input buffers")
	}
	return &ByteStream{
		pos:    0,
		sType:  MEMORY,
		binary: binary,
		cache:  make([]byte, CacheSize),
		r:      bytes.NewReader(binary),
	}, nil
}

// Len returns the current binary stream size.
func (bs *ByteStream) Len() int {
	return bs.r.Len()
}

// Pos returns the current position within the binary stream.
func (bs *ByteStream) Pos() int {
	return bs.pos
}

// Read len(b) bytes from the underlying stream.
func (bs *ByteStream) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = bs.r.Read(b)
	bs.pos += n
	return
}

// ReadAt len(b) bytes from the underlying stream at a given offset.
func (bs *ByteStream) ReadAt(b []byte, off int64) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = bs.r.ReadAt(b, off)
	bs.pos += n
	return
}

// Seek moves the stream iterator position to offset starting from whence.
func (bs *ByteStream) Seek(off int64, whence int) (n int64, err error) {
	n, err = bs.r.Seek(off, whence)
	if err != nil {
		return n, err
	}
	bs.pos += int(n)
	return n, err
}

// ReadUint8 reads a 1 byte little endian unsigned integer.
func (bs *ByteStream) ReadUint8(off int64) (uint8, error) {
	b := make([]byte, 1)
	n, err := bs.r.ReadAt(b, off)
	if n != 1 || err != nil {
		return 0, err
	}
	bs.pos += n
	return uint8(b[0]), err
}

// ReadUint16 reads a 2 byte little endian unsigned integer.
func (bs *ByteStream) ReadUint16(off int64) (uint16, error) {

	b := make([]byte, 2)
	n, err := bs.r.ReadAt(b, off)
	if n != 2 || err != nil {
		return 0, err
	}
	bs.pos += n
	return binary.LittleEndian.Uint16(b), err
}

// ReadUint32 reads a 4 byte little endian unsigned integer.
func (bs *ByteStream) ReadUint32(off int64) (uint32, error) {

	b := make([]byte, 4)
	n, err := bs.r.ReadAt(b, off)
	if n != 4 || err != nil {
		return 0, err
	}
	bs.pos += n
	return binary.LittleEndian.Uint32(b), err
}

// ReadUint64 reads a 8 byte little endian unsigned integer.
func (bs *ByteStream) ReadUint64(off int64) (uint64, error) {

	b := make([]byte, 8)
	n, err := bs.r.ReadAt(b, off)
	if n != 8 || err != nil {
		return 0, err
	}
	bs.pos += n
	return binary.LittleEndian.Uint64(b), err
}

// IsEOF checks if we reached the end of the bytestream.
func (bs *ByteStream) IsEOF() bool {
	// a bytestream wraps the bytereader interface
	// the Len method returns the number of unread bytes
	// in the underlying stream.
	return bs.r.Len() == 0
}

// Close underlying stream.
func (bs *ByteStream) Close() error {
	return nil
}
