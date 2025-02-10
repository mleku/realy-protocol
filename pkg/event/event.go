package event

import (
	"protocol.realy.lol/pkg/content"
	"protocol.realy.lol/pkg/decimal"
	"protocol.realy.lol/pkg/event/types"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/signature"
	"protocol.realy.lol/pkg/tags"
)

type E struct {
	Type      *types.T
	Pubkey    *pubkey.P
	Timestamp *decimal.T
	Tags      *tags.T
	Content   *content.C
	Signature *signature.S
}

// New creates a new event with some typical data already filled. This should be
// populated by some kind of editor.
//
// Simplest form of this would be to create a temporary file, open user's
// default editor with the event already populated, they enter the content
// field's message, and then after closing the editor it scans the text for e:
// and p: event and pubkey references and maybe #hashtags, updates the
// timestamp, and then signs it with a signing key, wraps in an event publish
// request, stamps and signs it and then pushes it to a configured relay
// address.
//
// Other more complex edit flows could be created but this one is for a simple
// flow as described. REALY events are text, and it is simple to make them
// literally edit as simple text files. REALY is natively text files, and the
// first composition client should just be a text editor.
func New(pk []byte, typ string) (ev *E, err error) {
	var p *pubkey.P
	p, err = pubkey.New(pk)
	ev = &E{
		Type:      types.New(typ),
		Pubkey:    p,
		Timestamp: decimal.Now(),
	}
	return
}

func (e *E) Marshal(d []byte) (r []byte, err error) {
	r = d
	if r, err = e.Type.Marshal(d); chk.E(err) {
		return
	}
	if r, err = e.Pubkey.Marshal(d); chk.E(err) {
		return
	}
	if r, err = e.Timestamp.Marshal(d); chk.E(err) {
		return
	}
	if r, err = e.Tags.Marshal(d); chk.E(err) {
		return
	}
	if r, err = e.Content.Marshal(d); chk.E(err) {
		return
	}
	if r, err = e.Signature.Marshal(d); chk.E(err) {
		return
	}
	return
}

func (e *E) Unmarshal(data []byte) (r []byte, err error) {

	return
}
