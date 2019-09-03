package inter

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/src/inter/pos"
)

type ApplyBlockFn func(block *Block, stateHash common.Hash, members pos.Members) (newStateHash common.Hash, mewMembers pos.Members)

// Block is a "chain" block.
type Block struct {
	Index      idx.Block
	Time       Timestamp
	Events     hash.Events
	SkippedTxs []uint // indexes of skipped txs, starting from first tx of first event, ending with last tx of last event

	PrevHash hash.Event

	Root    common.Hash
	Creator common.Address
}

func (b *Block) Hash() hash.Event {
	if len(b.Events) == 0 {
		return hash.ZeroEvent
	}
	return b.Events[len(b.Events)-1] // fiWitness is always a last event
}

// NewBlock makes block from topological ordered events.
func NewBlock(index idx.Block, time Timestamp, events hash.Events, prevHash hash.Event) *Block {
	return &Block{
		Index:      index,
		Time:       time,
		Events:     events,
		PrevHash:   prevHash,
		SkippedTxs: make([]uint, 0),
	}
}