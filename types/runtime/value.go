package runtime

import (
	"encoding/binary"
	"fmt"
)

type ValueType byte

const (
	ValueTypeI32 ValueType = 0x7F
	ValueTypeI64 ValueType = 0x7E
	ValueTypeF32 ValueType = 0x7D
	ValueTypeF64 ValueType = 0x7C
)

func (t ValueType) String() string {
	switch t {
	case ValueTypeI32:
		return "i32"
	case ValueTypeI64:
		return "i64"
	case ValueTypeF32:
		return "f32"
	case ValueTypeF64:
		return "f64"
	default:
		return "unknown"
	}
}

func ValueFrom(v any) (Value, error) {
	switch v := v.(type) {
	case int32:
		return ValueI32(v), nil
	case int64:
		return ValueI64(v), nil
	case bool:
		if v {
			return ValueI32(1), nil
		}
		return ValueI32(0), nil
	}
	return nil, fmt.Errorf("unsupported type %T", v)
}

type Value interface {
	Type() ValueType
}

type ValueI32 int32

func (ValueI32) Type() ValueType { return ValueTypeI32 }

type ValueI64 int64

func (ValueI64) Type() ValueType { return ValueTypeI64 }

type ValueF32 float32

func (ValueF32) Type() ValueType { return ValueTypeF32 }

type ValueF64 float64

func (ValueF64) Type() ValueType { return ValueTypeF64 }

func Falsy(v Value) bool {
	switch v := v.(type) {
	case ValueI32:
		return v == 0
	case ValueI64:
		return v == 0
	default:
		return false
	}
}

type LabelKind int

const (
	LabelKindBlock LabelKind = iota
	LabelKindLoop
	LabelKindIf
)

type Label struct {
	kind           LabelKind
	programCounter int
	stackPointer   int
	arity          int
}

func NewLabel(kind LabelKind, pc, sp, arity int) Label {
	return Label{
		kind:           kind,
		programCounter: pc,
		stackPointer:   sp,
		arity:          arity,
	}
}

func (l Label) Kind() LabelKind {
	return l.kind
}

func (l Label) ProgramCounter() int {
	return l.programCounter
}

func (l Label) StackPointer() int {
	return l.stackPointer
}

func (l Label) Arity() int {
	return l.arity
}

func writeValue(buf []byte, v Value) (int, error) {
	switch v := v.(type) {
	case ValueI32:
		return binary.Encode(buf, binary.LittleEndian, int32(v))
	case ValueI64:
		return binary.Encode(buf, binary.LittleEndian, int64(v))
	}
	return 0, fmt.Errorf("unsupported type %T", v)
}

func readValue[T Value](buf []byte) (int, T, error) {
	var v T
	n, err := binary.Decode(buf, binary.LittleEndian, &v)
	return n, v, err
}