package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

var endian = binary.LittleEndian

type ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func take[T ints](n T) func(r io.Reader) (io.Reader, error) {
	return func(r io.Reader) (io.Reader, error) {
		b := make([]byte, int(n))
		if _, err := io.ReadFull(r, b); err != nil {
			return nil, fmt.Errorf("failed to read %d bytes: %w", n, err)
		}
		return bytes.NewReader(b), nil
	}
}

func readByte(r io.Reader) (byte, error) {
	var (
		b [1]byte
	)
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, fmt.Errorf("failed to read byte: %w", err)
	}
	return b[0], nil
}

func readF32(r io.Reader) ([4]byte, error) {
	var b [4]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return b, err
	}
	return b, nil
}

func readF64(r io.Reader) ([8]byte, error) {
	var b [8]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return b, err
	}

	return b, nil
}
