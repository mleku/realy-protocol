package event

import (
	"protocol.realy.lol/pkg/content"
	"protocol.realy.lol/pkg/event/types"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/signature"
	"protocol.realy.lol/pkg/tags"
	"protocol.realy.lol/pkg/timestamp"
)

type Event struct {
	Type      *types.T
	Pubkey    *pubkey.P
	Timestamp *timestamp.T
	Tags      *tags.T
	Content   *content.C
	Signature *signature.S
}
