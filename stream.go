package binstream

// StreamType represents the stream type we're reading from (a File or Memory)
type StreamType int

const (
	// UNKNOWN is for unknown stream types.
	UNKNOWN StreamType = iota
	// FILE if we're reading from an on-disk file.
	FILE
	// MEMORY if we're reading from memory (byte slice)
	MEMORY
)

// Stream represents the generic interface each implementation should fullfill.
type Stream interface {
	Len() int // Current stream size
	// Small Reader Interface
	Read(b []byte) (n int, err error)              // Read len(b) from stream.
	ReadAt(b []byte, off int64) (n int, err error) // Read len(b) from stream starting at off.
	// Seek and position
	Pos() int                                        // Return current position in the stream
	Seek(off int64, whence int) (n int64, err error) // Move to offset by changing position
}
