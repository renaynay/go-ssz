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

// it would be serialized as:
//<proposer_index>
//<signed_header_1.message.slot>
//<signed_header_1.message.parent_root>
//<signed_header_1.message.state_root>
//<signed_header_1.message.body_root>
//<signed_header_1.signature>
//<signed_header_2.message.slot>
//<signed_header_2.message.parent_root>
//<signed_header_2.message.state_root>
//<signed_header_2.message.body_root>
//<signed_header_2.signature>
func MarshalProposerSlashing(ps *ethpb.ProposerSlashing) ([]byte, error) {
	bufLen := 0

	marshaledProposerIndex := encode.MarshalUint64(ps.ProposerIndex)
	bufLen += len(marshaledProposerIndex)

	marshaledHeader1, err := MarshalSignedBeaconBlockHeader(ps.Header_1)
	if err != nil {
		return []byte{}, err
	}
	bufLen += len(marshaledHeader1)

	marshaledHeader2, err := MarshalSignedBeaconBlockHeader(ps.Header_2)
	if err != nil {
		return []byte{}, err
	}
	bufLen += len(marshaledHeader2)

	// TODO are the other parts necessary? w/ prefix XXX

	out := make([]byte, bufLen)
	placeholder := 4

	copy(out[:placeholder], marshaledProposerIndex)
	placeholder += 4

	copy(out[placeholder:placeholder+len(marshaledHeader1)], marshaledHeader1) // TODO once you know length of marshaled header, then change the range declaration
	placeholder += len(marshaledHeader1)

	copy(out[placeholder:placeholder + len(marshaledHeader2)], marshaledHeader2)

	return out, nil
}

func MarshalSignedBeaconBlockHeader(sbbh *ethpb.SignedBeaconBlockHeader) ([]byte, error) {
	marshaledHeader, err := MarshalBeaconBlockHeader(sbbh.Header)
	if err != nil {
		return []byte{}, err
	}

	return append(marshaledHeader, sbbh.Signature...), nil
}

func MarshalBeaconBlockHeader(head *ethpb.BeaconBlockHeader) ([]byte, error) {
	marshaledSlot := encode.MarshalUint64(head.Slot)

	bufLen := (4 + len(head.ParentRoot) + len(head.StateRoot) + len(head.BodyRoot))
	out := make([]byte, bufLen)
	placeholder := 4

	copy(out[:placeholder], marshaledSlot)
	placeholder += 4

	copy(out[placeholder:placeholder+len(head.ParentRoot)], head.ParentRoot)
	placeholder += len(head.ParentRoot)

	copy(out[placeholder:placeholder+len(head.StateRoot)], head.StateRoot)
	placeholder += len(head.StateRoot)

	copy(out[placeholder:placeholder+len(head.BodyRoot)], head.BodyRoot)

	return out, nil
}

func MarshalBeaconBlockBody(bbb *ethpb.BeaconBlockBody) ([]byte, error) {
	return []byte{}, nil // TODO
}

