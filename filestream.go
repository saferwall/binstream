package binstream

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// ChunkSize is the default chunk size we use to read from files.
const ChunkSize = 1024

// FileStream represents an in-memory binary stream.
type FileStream struct {
	pos    int
	sType  StreamType
	cache  []byte
	binary []byte
	info   os.FileInfo
	f      *os.File
}

// Assert interface implementation checks
var _ Stream = (*FileStream)(nil)

// NewFileStream creates a new binary stream from a byte slice (can be memory mapped file).
// if chunk is set to true we read the first chunksize bytes.
func NewFileStream(filename string, chunk bool) (*FileStream, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	binary := make([]byte, ChunkSize)
	if chunk {
		n, err := f.Read(binary)
		if n != ChunkSize || err != nil {
			return nil, errors.New("failed to read from file")
		}
	}
	return &FileStream{
		pos:    0,
		sType:  FILE,
		cache:  make([]byte, CacheSize),
		binary: binary,
		info:   info,
		f:      f,
	}, nil
}

// Len returns the current binary stream size.
func (fs *FileStream) Len() int {
	return len(fs.binary)
}

// Pos returns the current position within the binary stream.
func (fs *FileStream) Pos() int {
	return fs.pos
}

// Read len(b) bytes from the underlying stream.
func (fs *FileStream) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = fs.f.Read(b)
	fs.pos += n
	return
}

// ReadAt len(b) bytes from the underlying stream at a given offset.
func (fs *FileStream) ReadAt(b []byte, off int64) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n, err = fs.f.ReadAt(b, off)
	fs.pos += n
	return
}

// Seek moves the stream iterator position to offset starting from whence.
func (fs *FileStream) Seek(off int64, whence int) (n int64, err error) {
	n, err = fs.f.Seek(off, whence)
	if err != nil {
		return n, err
	}
	fs.pos += int(n)
	return n, err
}

// ReadUint8 reads a 1 byte little endian unsigned integer.
func (fs *FileStream) ReadUint8(off int64) (uint8, error) {
	b := make([]byte, 1)
	n, err := fs.f.ReadAt(b, off)
	if n != 1 || err != nil {
		return 0, err
	}
	fs.pos += n
	return uint8(b[0]), err
}

// ReadUint16 reads a 2 byte little endian unsigned integer.
func (fs *FileStream) ReadUint16(off int64) (uint16, error) {

	b := make([]byte, 2)
	n, err := fs.f.ReadAt(b, off)
	if n != 2 || err != nil {
		return 0, err
	}
	fs.pos += n
	return binary.LittleEndian.Uint16(b), err
}

// ReadUint32 reads a 4 byte little endian unsigned integer.
func (fs *FileStream) ReadUint32(off int64) (uint32, error) {

	b := make([]byte, 4)
	n, err := fs.f.ReadAt(b, off)
	if n != 4 || err != nil {
		return 0, err
	}
	fs.pos += n
	return binary.LittleEndian.Uint32(b), err
}

// ReadUint64 reads a 8 byte little endian unsigned integer.
func (fs *FileStream) ReadUint64(off int64) (uint64, error) {

	b := make([]byte, 8)
	n, err := fs.f.ReadAt(b, off)
	if n != 8 || err != nil {
		return 0, err
	}
	fs.pos += n
	return binary.LittleEndian.Uint64(b), err
}

// IsEOF checks if we reached the end of the bytestream.
func (fs *FileStream) IsEOF() bool {
	// a bytestream wraps the bytereader interface
	// the Len method returns the number of unread bytes
	// in the underlying stream.
	n, err := fs.f.Read(nil)
	if n == 0 || err == io.EOF {
		return true
	}
	return false
}

// Close underlying filestream.
func (fs *FileStream) Close() error {
	return fs.f.Close()
}
