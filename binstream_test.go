package binstream

import (
	"bytes"
	"testing"
)

func TestBinaryStream(t *testing.T) {

	t.Run("TestNew", func(t *testing.T) {
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
			bs, err := New(tt.input)
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
			bs, err := New(tt.input)
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

}
