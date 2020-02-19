package decode

import (
	"encoding/binary"
	"errors"
)

func UnmarshalUint8(in byte) uint8 {
	return in
}

func UnmarshalUint32(in []byte) (uint32, error) {
	if len(in) != 4 {
		return uint32(0), errors.New("length other than 4 not allowed")
	}

	return binary.LittleEndian.Uint32(in), nil
}

func UnmarshalUint64(in []byte) (uint64, error) {
	if len(in) != 8 {
		return uint64(0), errors.New("length other than 8 not allowed")
	}

	return binary.LittleEndian.Uint64(in), nil
}