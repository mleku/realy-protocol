package event

import (
	"protocol.realy.lol/pkg/content"
	"protocol.realy.lol/pkg/event/types"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/signature"
)

type Event struct {
	Type      *types.T
	Pubkey    *pubkey.P
	Timestamp int64
	Tags      [][]byte
	Content   *content.C
	Signature *signature.S
}
