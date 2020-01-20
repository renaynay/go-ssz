package types

import "testing"

func TestMarshalUnmarshal(t *testing.T) {
	array := Bytes32Array{'A', 'B', 'C', 'D'}
	junk := make([]byte, 32)
	array.MarshalTo(junk)
	var newarray Bytes32Array
	newarray.Unmarshal(junk)

	if newarray != array {
		t.Fatalf("wanted %v but got %v", array, newarray)
	}
}
