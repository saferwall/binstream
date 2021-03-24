package binstream

import (
	"bytes"
	"io"
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
			fs, err := NewFileStream(tt.input)
			if fs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			fs.Close()
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
			fs, err := NewFileStream(tt.input)
			if fs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := fs.Read(b)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.expected) {
				t.Fatalf("copied bytes from BinaryStream are not consistent with input expected %v got %v", tt.expected, b)
			}
			fs.Close()
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
			fs, err := NewFileStream(tt.input)
			if fs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 4)
			n, err := fs.ReadAt(b, 4)
			if n != len(b) || err != nil {
				t.Fatal("failed to read from BinaryStream with error", err)
			}
			if !bytes.Equal(b, tt.expected) {
				t.Fatalf("copied bytes from BinaryStream are not consistent with input expected %v got %v", tt.expected, b)
			}
			fs.Close()

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
				expected: nil,
				err:      io.EOF,
			},
		}

		for _, tt := range testCases {
			fs, err := NewFileStream(tt.input)
			if fs == nil || err != nil {
				t.Fatal("failed to create new BinaryStream instance with error", err)
			}
			b := make([]byte, 1)
			off := int64(142144) // size of /bin/ls
			_, err = fs.ReadAt(b, off)
			if err != tt.err {
				t.Fatal("expected EOF but got ", err)
			}
			fs.Close()

		}
	})

}
