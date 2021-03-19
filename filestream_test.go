package binstream

import (
	"bytes"
	"testing"
)

func TestFileStream(t *testing.T) {

	t.Run("TestNewFileStream", func(t *testing.T) {
		testCases := []struct {
			input string
			err   error
		}{
			{
				input: "/bin/ls",
				err:   nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewFileStream(tt.input, true)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			bs.Close()
		}
	})
	t.Run("TestRead", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected []byte
			err      error
		}{
			{
				input:    "/bin/ls",
				expected: []byte("\177ELF"),
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewFileStream(tt.input, false)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := bs.Read(b)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.expected) {
				t.Fatalf("copied bytes from BinaryStream are not consistent with input expected %v got %v", tt.expected, b)
			}
			bs.Close()
		}
	})
	t.Run("TestReadAt", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected []byte
			err      error
		}{
			{
				input:    "/bin/ls",
				expected: []byte{0x2, 0x1, 0x1, 0x0},
				err:      nil,
			},
		}

		for _, tt := range testCases {
			bs, err := NewFileStream(tt.input, true)
			if bs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := bs.ReadAt(b, 4)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.expected) {
				t.Fatalf("copied bytes from BinaryStream are not consistent with input expected %v got %v", tt.expected, b)
			}
			bs.Close()

		}
	})

}
