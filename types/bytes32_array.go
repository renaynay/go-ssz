package types

// Bytes32Array defines a compatible gogo protobuf type.
type Bytes32Array [32]byte

// Size of the array.
func (b *Bytes32Array) Size() int {
	return len(b)
}

// MarshalTo serializes the array to the buffer.
func (b *Bytes32Array) MarshalTo(data []byte) (int, error) {
	return copy(data[:b.Size()], b[:]), nil
}

// Unmarshal deserializes into the provided array type.
func (b *Bytes32Array) Unmarshal(data []byte) error {
	copy(b[:], data)
	return nil
}
