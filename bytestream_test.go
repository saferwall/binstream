package binstream

import (
	"bytes"
	"testing"
)

func TestBinaryStream(t *testing.T) {

	t.Run("TestNewByteStream", func(t *testing.T) {
		testCases := []struct {
			input []byte
			err   error
		}{
			{
				input: []byte{0xff, 0xff, 0xff, 0xff},
				err:   nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
		}
	})
	t.Run("TestRead", func(t *testing.T) {
		testCases := []struct {
			input []byte
			err   error
		}{
			{
				input: []byte{0xff, 0xff, 0xff, 0xff},
				err:   nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := bs.Read(b)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.input) {
				t.Fatal("copied bytes from BinaryStream are not consistent with input")
			}
		}
	})
	t.Run("TestReadAt", func(t *testing.T) {
		testCases := []struct {
			input []byte
			err   error
		}{
			{
				input: []byte{0xff, 0xff, 0xab, 0xab},
				err:   nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 2)
			n, err := bs.ReadAt(b, 2)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.input[2:]) {
				t.Fatal("copied bytes from BinaryStream are not consistent with input")
			}
		}
	})
	t.Run("TestReadUint8", func(t *testing.T) {
		testCases := []struct {
			input    []byte
			expected []uint8
			err      error
		}{
			{
				input:    []byte{0xff, 0xff, 0xa, 0xb},
				expected: []uint8{0xff, 0xa},
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			n, err := bs.ReadUint8(0)
			if n != tt.expected[0] || err != nil {
				t.Fatalf("failed to read from BinaryStream with error %v : expected %d got %d", err, tt.expected[0], n)
			}
			n, err = bs.ReadUint8(2)
			if n != tt.expected[1] || err != nil {
				t.Fatalf("failed to read from BinaryStream with error %v : expected %d got %d", err, tt.expected[1], n)
			}
		}
	})

	t.Run("TestReadUint16", func(t *testing.T) {
		testCases := []struct {
			input    []byte
			expected []uint16
			err      error
		}{
			{
				input:    []byte{0xe8, 0x03, 0xd0, 0x07},
				expected: []uint16{0x03e8, 0x07d0},
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			n, err := bs.ReadUint16(0)
			if n != tt.expected[0] || err != nil {
				t.Fatalf("failed to read from BinaryStream with error %v : expected %d got %d", err, tt.expected[0], n)
			}
			n, err = bs.ReadUint16(2)
			if n != tt.expected[1] || err != nil {
				t.Fatalf("failed to read from BinaryStream with error %v : expected %d got %d", err, tt.expected[1], n)
			}
		}
	})
	t.Run("TestReadUint32", func(t *testing.T) {
		testCases := []struct {
			input    []byte
			expected []uint32
			err      error
		}{
			{
				input:    []byte{0xe8, 0x03, 0xd0, 0x07},
				expected: []uint32{0x07d003e8},
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			n, err := bs.ReadUint32(0)
			if n != tt.expected[0] || err != nil {
				t.Fatalf("failed to read from BinaryStream with error %v : expected %d got %d", err, tt.expected[0], n)
			}
		}
	})
	t.Run("TestIsEOF", func(t *testing.T) {
		testCases := []struct {
			input    []byte
			expected bool
			err      error
		}{
			{
				input:    []byte{0xe8, 0x03, 0xd0, 0x07},
				expected: true,
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewByteStream(tt.input)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := bs.Read(b)
			if n != 4 || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if bs.IsEOF() != tt.expected {
				t.Fatal("failed to read entire stream expected EOF got none")
			}
		}
	})

}
