package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
)

type Module struct {
	magic   string
	version uint32
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func decode(r io.Reader) (*Module, error) {
	magic, version, err := decodePreamble(r)
	if err != nil {
		return nil, err
	}

	for {
		code, size, err := decodeSectionHeader(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to decode section header: %w", err)
		}

		section := make([]byte, size)
		if _, err := io.ReadFull(r, section); err != nil {
			return nil, fmt.Errorf("failed to read section: %w", err)
		}

		switch code {
		default:
			return nil, fmt.Errorf("unsupported section code: %d", code)
		}
	}

	return &Module{
		magic:   magic,
		version: version,
	}, nil
}

func decodePreamble(r io.Reader) (string, uint32, error) {
	var (
		magic   [4]byte
		version [4]byte
	)
	if _, err := io.ReadFull(r, magic[:]); err != nil {
		return "", 0, fmt.Errorf("failed to read magic binary: %w", err)
	}
	if string(magic[:]) != "\x00asm" {
		return "", 0, fmt.Errorf("invalid magic header: %x", magic[:])
	}
	if _, err := io.ReadFull(r, version[:]); err != nil {
		return "", 0, fmt.Errorf("failed to read version: %w", err)
	}

	return string(magic[:]), binary.LittleEndian.Uint32(version[:]), nil
}

func decodeSectionHeader(r io.Reader) (SectionCode, uint32, error) {
	var (
		code [1]byte
	)
	if _, err := io.ReadFull(r, code[:]); err != nil {
		return 0, 0, fmt.Errorf("failed to read section code: %w", err)
	}

	size, err := leb128.Uint32(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read section size: %w", err)
	}

	return SectionCode(code[0]), size, nil
}