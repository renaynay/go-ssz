// Code generated by yaml_to_go. DO NOT EDIT.
// source: ssz_mainnet_random.yaml

package autogenerated

import (
	"github.com/prysmaticlabs/go-bitfield"
)

type MainnetFork struct {
	PreviousVersion []byte `json:"previous_version" ssz-size:"4"`
	CurrentVersion  []byte `json:"current_version" ssz-size:"4"`
	Epoch           uint64 `json:"epoch"`
}

type MainnetCheckpoint struct {
	Epoch uint64 `json:"epoch"`
	Root  []byte `json:"root" ssz-size:"32"`
}

type MainnetValidator struct {
	Pubkey                     []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz-size:"32"`
	EffectiveBalance           uint64 `json:"effective_balance"`
	Slashed                    bool   `json:"slashed"`
	ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
	ActivationEpoch            uint64 `json:"activation_epoch"`
	ExitEpoch                  uint64 `json:"exit_epoch"`
	WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
}

type MainnetCrosslink struct {
	Shard      uint64 `json:"shard"`
	ParentRoot []byte `json:"parent_root" ssz-size:"32"`
	StartEpoch uint64 `json:"start_epoch"`
	EndEpoch   uint64 `json:"end_epoch"`
	DataRoot   []byte `json:"data_root" ssz-size:"32"`
}

type MainnetAttestationData struct {
	BeaconBlockRoot []byte            `json:"beacon_block_root" ssz-size:"32"`
	Source          MainnetCheckpoint `json:"source"`
	Target          MainnetCheckpoint `json:"target"`
	Crosslink       MainnetCrosslink  `json:"crosslink"`
}

type MainnetAttestationAndCustodyBit struct {
	Data       MainnetAttestationData `json:"data"`
	CustodyBit bool                   `json:"custody_bit"`
}

type MainnetIndexedAttestation struct {
	CustodyBit0Indices []uint64               `json:"custody_bit_0_indices" ssz-max:"4096"`
	CustodyBit1Indices []uint64               `json:"custody_bit_1_indices" ssz-max:"4096"`
	Data               MainnetAttestationData `json:"data"`
	Signature          []byte                 `json:"signature" ssz-size:"96"`
}

type MainnetPendingAttestation struct {
	AggregationBits bitfield.Bitlist       `json:"aggregation_bits" ssz-max:"4096"`
	Data            MainnetAttestationData `json:"data"`
	InclusionDelay  uint64                 `json:"inclusion_delay"`
	ProposerIndex   uint64                 `json:"proposer_index"`
}

type MainnetEth1Data struct {
	DepositRoot  []byte `json:"deposit_root" ssz-size:"32"`
	DepositCount uint64 `json:"deposit_count"`
	BlockHash    []byte `json:"block_hash" ssz-size:"32"`
}

type MainnetHistoricalBatch struct {
	BlockRoots [][]byte `json:"block_roots" ssz-size:"8192,32"`
	StateRoots [][]byte `json:"state_roots" ssz-size:"8192,32"`
}

type MainnetDepositData struct {
	Pubkey                []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials []byte `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64 `json:"amount"`
	Signature             []byte `json:"signature" ssz-size:"96"`
}

type MainnetCompactCommittee struct {
	Pubkeys           [][]byte `json:"pubkeys" ssz-size:"?,48" ssz-max:"4096"`
	CompactValidators []uint64 `json:"compact_validators" ssz-max:"4096"`
}

type MainnetBlockHeader struct {
	Slot       uint64 `json:"slot"`
	ParentRoot []byte `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte `json:"state_root" ssz-size:"32"`
	BodyRoot   []byte `json:"body_root" ssz-size:"32"`
	Signature  []byte `json:"signature" ssz-size:"96"`
}

type MainnetProposerSlashing struct {
	ProposerIndex uint64             `json:"proposer_index"`
	Header1       MainnetBlockHeader `json:"header_1"`
	Header2       MainnetBlockHeader `json:"header_2"`
}

type MainnetAttesterSlashing struct {
	Attestation1 MainnetIndexedAttestation `json:"attestation_1"`
	Attestation2 MainnetIndexedAttestation `json:"attestation_2"`
}

type MainnetAttestation struct {
	AggregationBits bitfield.Bitlist       `json:"aggregation_bits" ssz-max:"4096"`
	Data            MainnetAttestationData `json:"data"`
	CustodyBits     bitfield.Bitlist       `json:"custody_bits" ssz-max:"4096"`
	Signature       []byte                 `json:"signature" ssz-size:"96"`
}

type MainnetDeposit struct {
	Proof [][]byte           `json:"proof" ssz-size:"33,32"`
	Data  MainnetDepositData `json:"data"`
}

type MainnetVoluntaryExit struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
	Signature      []byte `json:"signature" ssz-size:"96"`
}

type MainnetTransfer struct {
	Sender    uint64 `json:"sender"`
	Recipient uint64 `json:"recipient"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
	Slot      uint64 `json:"slot"`
	Pubkey    []byte `json:"pubkey" ssz-size:"48"`
	Signature []byte `json:"signature" ssz-size:"96"`
}

type MainnetBlockBody struct {
	RandaoReveal      []byte                    `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          MainnetEth1Data           `json:"eth1_data"`
	Graffiti          []byte                    `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []MainnetProposerSlashing `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []MainnetAttesterSlashing `json:"attester_slashings" ssz-max:"1"`
	Attestations      []MainnetAttestation      `json:"attestations" ssz-max:"128"`
	Deposits          []MainnetDeposit          `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []MainnetVoluntaryExit    `json:"voluntary_exits" ssz-max:"16"`
	Transfers         []MainnetTransfer         `json:"transfers" ssz-max:"0"`
}

type MainnetBlock struct {
	Slot       uint64           `json:"slot"`
	ParentRoot []byte           `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte           `json:"state_root" ssz-size:"32"`
	Body       MainnetBlockBody `json:"body"`
	Signature  []byte           `json:"signature" ssz-size:"96"`
}

type MainnetBeaconState struct {
	GenesisTime            uint64             `json:"genesis_time"`
	Slot                   uint64             `json:"slot"`
	Fork                   MainnetFork        `json:"fork"`
	LatestBlockHeader      MainnetBlockHeader `json:"latest_block_header"`
	BlockRoots             [][]byte           `json:"block_roots" ssz-size:"8192,32"`
	StateRoots             [][]byte           `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots        [][]byte           `json:"historical_roots" ssz-size:"?,32" ssz-max:"16777216"`
	Eth1Data               MainnetEth1Data    `json:"eth1_data"`
	Eth1DataVotes          []MainnetEth1Data  `json:"eth1_data_votes" ssz-max:"1024"`
	Eth1DepositIndex       uint64             `json:"eth1_deposit_index"`
	Validators             []MainnetValidator `json:"validators" ssz-max:"1099511627776"`
	Balances               []uint64           `json:"balances" ssz-max:"1099511627776"`
	StartShard             uint64             `json:"start_shard"`
	RandaoMixes            [][]byte           `json:"randao_mixes" ssz-size:"65536,32"`
	ActiveIndexRoots       [][]byte           `json:"active_index_roots" ssz-size:"65536,32"`
	CompactCommitteesRoots [][]byte           `json:"compact_committees_roots" ssz-size:"65536,32"`
	Slashings              []uint64           `json:"slashings" ssz-size:"8192"`

	PreviousEpochAttestations []MainnetPendingAttestation `json:"previous_epoch_attestations" ssz-max:"8192"`
	CurrentEpochAttestations  []MainnetPendingAttestation `json:"current_epoch_attestations" ssz-max:"8192"`
	PreviousCrosslinks        []MainnetCrosslink          `json:"previous_crosslinks" ssz-size:"1024"`
	CurrentCrosslinks         []MainnetCrosslink          `json:"current_crosslinks" ssz-size:"1024"`
	JustificationBits         bitfield.Bitvector4         `json:"justification_bits" ssz-size:"1"`

	PreviousJustifiedCheckpoint MainnetCheckpoint `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  MainnetCheckpoint `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         MainnetCheckpoint `json:"finalized_checkpoint"`
}

type SszMainnetTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Attestation struct {
			Value       MainnetAttestation `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"Attestation,omitempty"`
		AttestationData struct {
			Value      MainnetAttestationData `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"AttestationData,omitempty"`
		AttestationDataAndCustodyBit struct {
			Value      MainnetAttestationAndCustodyBit `json:"value"`
			Serialized []byte                          `json:"serialized"`
			Root       []byte                          `json:"root" ssz:"size=32"`
		} `json:"AttestationDataAndCustodyBit,omitempty"`
		AttesterSlashing struct {
			Value      MainnetAttesterSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"AttesterSlashing,omitempty"`
		BeaconBlock struct {
			Value       MainnetBlock `json:"value"`
			Serialized  []byte       `json:"serialized"`
			Root        []byte       `json:"root" ssz:"size=32"`
			SigningRoot []byte       `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlock,omitempty"`
		BeaconBlockBody struct {
			Value      MainnetBlockBody `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"BeaconBlockBody,omitempty"`
		BeaconBlockHeader struct {
			Value       MainnetBlockHeader `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlockHeader,omitempty"`
		BeaconState struct {
			Value      MainnetBeaconState `json:"value"`
			Serialized []byte             `json:"serialized"`
			Root       []byte             `json:"root" ssz:"size=32"`
		} `json:"BeaconState,omitempty"`
		Checkpoint struct {
			Value      MainnetCheckpoint `json:"value"`
			Serialized []byte            `json:"serialized"`
			Root       []byte            `json:"root" ssz:"size=32"`
		} `json:"Checkpoint,omitempty"`
		CompactCommittee struct {
			Value      MainnetCompactCommittee `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"CompactCommittee,omitempty"`
		Crosslink struct {
			Value      MainnetCrosslink `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Crosslink,omitempty"`
		Deposit struct {
			Value      MainnetDeposit `json:"value"`
			Serialized []byte         `json:"serialized"`
			Root       []byte         `json:"root" ssz:"size=32"`
		} `json:"Deposit,omitempty"`
		DepositData struct {
			Value       MainnetDepositData `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"DepositData,omitempty"`
		Eth1Data struct {
			Value      MainnetEth1Data `json:"value"`
			Serialized []byte          `json:"serialized"`
			Root       []byte          `json:"root" ssz:"size=32"`
		} `json:"Eth1Data,omitempty"`
		Fork struct {
			Value      MainnetFork `json:"value"`
			Serialized []byte      `json:"serialized"`
			Root       []byte      `json:"root" ssz:"size=32"`
		} `json:"Fork,omitempty"`
		HistoricalBatch struct {
			Value      MainnetHistoricalBatch `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"HistoricalBatch,omitempty"`
		IndexedAttestation struct {
			Value       MainnetIndexedAttestation `json:"value"`
			Serialized  []byte                    `json:"serialized"`
			Root        []byte                    `json:"root" ssz:"size=32"`
			SigningRoot []byte                    `json:"signing_root" ssz:"size=32"`
		} `json:"IndexedAttestation,omitempty"`
		PendingAttestation struct {
			Value      MainnetPendingAttestation `json:"value"`
			Serialized []byte                    `json:"serialized"`
			Root       []byte                    `json:"root" ssz:"size=32"`
		} `json:"PendingAttestation,omitempty"`
		ProposerSlashing struct {
			Value      MainnetProposerSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"ProposerSlashing,omitempty"`
		Transfer struct {
			Value       MainnetTransfer `json:"value"`
			Serialized  []byte          `json:"serialized"`
			Root        []byte          `json:"root" ssz:"size=32"`
			SigningRoot []byte          `json:"signing_root" ssz:"size=32"`
		} `json:"Transfer,omitempty"`
		Validator struct {
			Value      MainnetValidator `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Validator,omitempty"`
		VoluntaryExit struct {
			Value       MainnetVoluntaryExit `json:"value"`
			Serialized  []byte               `json:"serialized"`
			Root        []byte               `json:"root" ssz:"size=32"`
			SigningRoot []byte               `json:"signing_root" ssz:"size=32"`
		} `json:"VoluntaryExit,omitempty"`
	} `json:"test_cases"`
}
