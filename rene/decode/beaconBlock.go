package decode

import (
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/renaynay/go-ssz/rene/basic-types/decode"
)

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
func UnmarshalProposerSlashing(in []byte, ps *ethpb.ProposerSlashing) error {
	rawProposerIndex := in[:8]
	proposerIndex, err := decode.UnmarshalUint64(rawProposerIndex)
	if err != nil {
		return err
	}
	ps.ProposerIndex = proposerIndex

	headerLen := (len(in) - 8) / 2 // TODO how to know the length?

	rawHeader1 := in[8:8+headerLen] // TODO how to know the length?
	var header1 ethpb.SignedBeaconBlockHeader

	err = UnmarshalSignedBeaconBlockHeader(rawHeader1, &header1)
	if err != nil {
		return err
	}
	ps.Header_1 = &header1

	rawHeader2 := in[8+headerLen:] // TODO how to know the length?
	var header2 ethpb.SignedBeaconBlockHeader

	err = UnmarshalSignedBeaconBlockHeader(rawHeader2, &header2)
	if err != nil {
		return err
	}
	ps.Header_2 = &header2

	return nil
}

func UnmarshalSignedBeaconBlockHeader(in []byte, sbbh *ethpb.SignedBeaconBlockHeader) error {

	return nil
}

func UnmarshalBeaconBlockHeader(in []byte, bbh *ethpb.BeaconBlockHeader) error {

	return nil
}