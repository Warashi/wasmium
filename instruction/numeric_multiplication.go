package instruction

import (
	"io"

	"github.com/Warashi/wasmium/opcode"
)

type I32Mul struct{}

func (i *I32Mul) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Mul
}

func (i *I32Mul) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Mul struct{}

func (i *I64Mul) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Mul
}

func (i *I64Mul) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Mul struct{}

func (f *F32Mul) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Mul
}

func (f *F32Mul) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Mul struct{}

func (f *F64Mul) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Mul
}

func (f *F64Mul) ReadOperandsFrom(r io.Reader) error {
	return nil
}
