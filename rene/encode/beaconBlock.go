package encode

import (
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/renaynay/go-ssz/rene/basic-types/encode"
)

func MarshalBeaconBlock(bb *ethpb.BeaconBlock) ([]byte, error) {
	out := make([]byte, 0)
	offset := uint64(0)

	marshaledSlot := encode.MarshalUint64(bb.Slot)
	out = append(out, marshaledSlot...)
	offset++

	marshaledParentRootLen := encode.MarshalUint32(uint32(len(bb.ParentRoot)))
	out = append(out, marshaledParentRootLen...)
	out = append(out, bb.ParentRoot...)
	offset += uint64(len(marshaledParentRootLen) + len(bb.ParentRoot))

	marshaledStateRootLen := encode.MarshalUint32(uint32(len(bb.StateRoot)))
	out = append(out, marshaledStateRootLen...)
	out = append(out, bb.StateRoot...)
	offset += uint64(len(marshaledStateRootLen) + len(bb.StateRoot))

	marshaledBody, err := MarshalBeaconBlockBody(bb.Body)
	if err != nil {
		return nil, err
	}
	out = append(out, marshaledBody...)
	offset += uint64(len(marshaledBody))

	marshaledOffset := encode.MarshalUint64(offset)
	marshaledOffset = append(marshaledOffset, out...)

	return marshaledOffset, nil
}

func MarshalBeaconBlockBody(bbb *ethpb.BeaconBlockBody) ([]byte, error) {

}