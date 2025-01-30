package event

import (
	"protocol.realy.lol/pkg/event/types"
)

type Event struct {
	Type      types.T
	Pubkey    []byte
	Timestamp int64
	Tags      [][]byte
	Content   []byte
	Signature []byte
}
