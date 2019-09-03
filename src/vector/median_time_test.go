package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/inter/pos"
	"github.com/Fantom-foundation/go-lachesis/src/kvdb/memorydb"
	"github.com/Fantom-foundation/go-lachesis/src/logger"
)

func testMedianTime(t *testing.T, dag string, weights []pos.Stake, claimedTimes map[string]inter.Timestamp, medianTimes map[string]inter.Timestamp, genesis inter.Timestamp) {
	logger.SetTestMode(t)
	assertar := assert.New(t)

	ordered := make([]*inter.Event, 0)
	peers, _, named := inter.ASCIIschemeForEach(dag, inter.ForEachEvent{
		Build: func(e *inter.Event, name string) *inter.Event {
			e.ClaimedTime = claimedTimes[name]
			return e
		},
		Process: func(e *inter.Event, name string) {
			ordered = append(ordered, e)
		},
	})

	members := make(pos.Members, len(peers))
	for i, peer := range peers {
		members.Set(peer, weights[i])
	}

	vi := NewIndex(members, memorydb.New())

	// push
	for _, e := range ordered {
		vi.Add(e)
		vi.Flush()
	}

	// check
	for name, e := range named {
		expected, ok := medianTimes[name]
		if !ok {
			continue
		}
		assertar.Equal(expected, vi.MedianTime(e.Hash(), genesis), name)
	}
}

func TestMedianTimeAscii(t *testing.T) {
	dagStr := `
 ║
 nodeA001
 ║
 nodeA012
 ║            ║
 ║            nodeB001
 ║            ║            ║
 ║            ╠═══════════ nodeC001
 ║║           ║            ║            ║
 ║╚══════════─╫─══════════─╫─══════════ nodeD001
║║            ║            ║            ║
╚ nodeA002════╬════════════╬════════════╣
 ║║           ║            ║            ║
 ║╚══════════─╫─══════════─╫─══════════ nodeD002
 ║            ║            ║            ║
 nodeA003════─╫─══════════─╫─═══════════╣
 ║            ║            ║
 ╠════════════nodeB002     ║
 ║                         ║
 ╠═════════════════════════nodeC002
`

	weights := []pos.Stake{3, 4, 2, 1}
	genesisTime := inter.Timestamp(1)
	claimedTimes := map[string]inter.Timestamp{
		"nodeA001": inter.Timestamp(111),
		"nodeB001": inter.Timestamp(112),
		"nodeC001": inter.Timestamp(13),
		"nodeD001": inter.Timestamp(14),
		"nodeA002": inter.Timestamp(120),
		"nodeD002": inter.Timestamp(20),
		"nodeA012": inter.Timestamp(120),
		"nodeA003": inter.Timestamp(131),
		"nodeB002": inter.Timestamp(124),
		"nodeC002": inter.Timestamp(20),
	}
	medianTimes := map[string]inter.Timestamp{
		"nodeA001": genesisTime,
		"nodeB001": genesisTime,
		"nodeC001": inter.Timestamp(13),
		"nodeD001": genesisTime,
		"nodeA002": inter.Timestamp(112),
		"nodeD002": genesisTime,
		"nodeA012": genesisTime,
		"nodeA003": inter.Timestamp(112),
		"nodeB002": inter.Timestamp(124),
		"nodeC002": inter.Timestamp(20),
	}
	t.Run("medianTimeWithForks", func(t *testing.T) {
		testMedianTime(t, dagStr, weights, claimedTimes, medianTimes, genesisTime)
	})
}

func TestMedianTime(t *testing.T) {
	peers := inter.GenNodes(5)
	members := make(pos.Members, len(peers))

	weights := []pos.Stake{5, 4, 3, 2, 1}
	for i, peer := range peers {
		members.Set(peer, weights[i])
	}

	vi := NewIndex(members, memorydb.New())

	assertar := assert.New(t)
	{ // seq=0
		e := &event{
			EventHeaderData: &inter.NewEvent().EventHeaderData,
		}
		// member indexes are sorted by stake amount
		e.HighestBefore = []HighestBefore{
			{
				Seq:         0,
				ClaimedTime: 100,
			},
			{
				Seq:         0,
				ClaimedTime: 100,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
		}
		vi.SetEvent(e)
		assertar.Equal(inter.Timestamp(1), vi.MedianTime(e.Hash(), 1))
	}

	{ // fork seen = true
		e := &event{
			EventHeaderData: &inter.NewEvent().EventHeaderData,
		}
		// member indexes are sorted by stake amount
		e.HighestBefore = []HighestBefore{
			{
				Seq:         0,
				ClaimedTime: 100,
				IsForkSeen:  true,
			},
			{
				Seq:         0,
				ClaimedTime: 100,
				IsForkSeen:  true,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
			{
				Seq:         1,
				ClaimedTime: 10,
			},
		}
		vi.SetEvent(e)
		assertar.Equal(inter.Timestamp(10), vi.MedianTime(e.Hash(), 1))
	}

	{ // normal
		e := &event{
			EventHeaderData: &inter.NewEvent().EventHeaderData,
		}
		// member indexes are sorted by stake amount
		e.HighestBefore = []HighestBefore{
			{
				Seq:         1,
				ClaimedTime: 11,
			},
			{
				Seq:         2,
				ClaimedTime: 12,
			},
			{
				Seq:         2,
				ClaimedTime: 13,
			},
			{
				Seq:         3,
				ClaimedTime: 14,
			},
			{
				Seq:         4,
				ClaimedTime: 15,
			},
		}
		vi.SetEvent(e)
		assertar.Equal(inter.Timestamp(12), vi.MedianTime(e.Hash(), 1))
	}

}