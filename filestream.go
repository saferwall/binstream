package binstream

import (
	"io"
	"os"

	"github.com/edsrzf/mmap-go"
)

// ChunkSize is the default chunk size we use to read from files.
const ChunkSize = 1024

// FileStream represents an in-memory binary stream.
type FileStream struct {
	sType StreamType
	info  os.FileInfo
	f     *os.File
	bs    *ByteStream
	data  mmap.MMap
}

// Assert interface implementation checks
var _ Stream = (*FileStream)(nil)

// NewFileStream creates a new binary stream from a memory mapped file).
func NewFileStream(filename string) (*FileStream, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}
	// Memory map the file insead of using read/write.
	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		f.Close()
		return nil, err
	}
	bs, err := NewByteStream(data)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &FileStream{
		sType: FILE,
		info:  info,
		f:     f,
		bs:    bs,
		data:  data,
	}, nil
}

// Len returns the current binary stream size.
func (fs *FileStream) Len() int {
	return fs.bs.Len()
}

// Pos returns the current position within the binary stream.
func (fs *FileStream) Pos() int {
	return fs.bs.pos
}

// Read len(b) bytes from the underlying stream.
func (fs *FileStream) Read(b []byte) (n int, err error) {
	return fs.bs.Read(b)
}

// ReadAt len(b) bytes from the underlying stream at a given offset.
func (fs *FileStream) ReadAt(b []byte, off int64) (n int, err error) {
	return fs.bs.ReadAt(b, off)
}

// Seek moves the stream iterator position to offset starting from whence.
func (fs *FileStream) Seek(off int64, whence int) (n int64, err error) {
	return fs.bs.Seek(off, whence)
}

// ReadUint8 reads a 1 byte little endian unsigned integer.
func (fs *FileStream) ReadUint8(off int64) (uint8, error) {
	return fs.bs.ReadUint8(off)
}

// ReadUint16 reads a 2 byte little endian unsigned integer.
func (fs *FileStream) ReadUint16(off int64) (uint16, error) {
	return fs.bs.ReadUint16(off)
}

// ReadUint32 reads a 4 byte little endian unsigned integer.
func (fs *FileStream) ReadUint32(off int64) (uint32, error) {
	return fs.bs.ReadUint32(off)
}

// ReadUint64 reads a 8 byte little endian unsigned integer.
func (fs *FileStream) ReadUint64(off int64) (uint64, error) {

	return fs.bs.ReadUint64(off)
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
	if fs.data != nil {
		_ = fs.data.Unmap()
	}
	return fs.f.Close()
}
