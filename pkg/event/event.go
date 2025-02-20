package event

import (
	"realy.lol/sha256"
	"realy.lol/signer"

	"protocol.realy.lol/pkg/content"
	"protocol.realy.lol/pkg/decimal"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/separator"
	"protocol.realy.lol/pkg/signature"
	"protocol.realy.lol/pkg/tags"
	"protocol.realy.lol/pkg/types"
)

type E struct {
	id        []byte
	Type      *types.T
	Pubkey    *pubkey.P
	Timestamp *decimal.T
	Tags      *tags.T
	Content   *content.C
	Signature *signature.S
	encoded   []byte
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

// Invalidate empties the existing encoded cache of the event. This needs to be
// called in case of mutating its fields. It also nils the signature.
func (e *E) Invalidate() { e.encoded = e.encoded[:0]; e.Signature = nil; e.id = nil }

func (e *E) Sign(s signer.I) (err error) {
	var h []byte
	if h, err = e.Hash(); chk.E(err) {
		return
	}
	var sig []byte
	if sig, err = s.Sign(h); chk.E(err) {
		return
	}
	if e.Signature, err = signature.New(sig); chk.E(err) {
		return
	}
	return
}

func (e *E) Encode(d []byte) (r []byte, err error) {
	r = d
	if e.Type == nil {
		err = errorf.E("type is not defined for event")
		return
	}
	if r, err = e.Type.Marshal(r); chk.E(err) {
		return
	}
	r = separator.Add(r)
	if e.Pubkey == nil {
		err = errorf.E("pubkey is not defined for event")
		return
	}
	// log.I.S(r)
	if r, err = e.Pubkey.Marshal(r); chk.E(err) {
		return
	}
	r = separator.Add(r)
	if e.Timestamp == nil {
		err = errorf.E("timestamp is not defined for event")
		return
	}
	if r, err = e.Timestamp.Marshal(r); chk.E(err) {
		return
	}
	r = separator.Add(r)
	if r, err = e.Tags.Marshal(r); chk.E(err) {
		return
	}
	if e.Content != nil {
		if r, err = e.Content.Marshal(r); chk.E(err) {
			return
		}
		r = separator.Add(r)
	}
	e.encoded = r
	return
}

func (e *E) Hash() (h []byte, err error) {
	var b []byte
	if e.encoded == nil {
		if e.encoded, err = e.Encode(nil); chk.E(err) {
			return
		}
		b = e.encoded
	}
	hh := sha256.Sum256(b)
	h = hh[:]
	e.id = h
	return
}

func (e *E) Marshal(d []byte) (r []byte, err error) {
	if r, err = e.Encode(d); chk.E(err) {
		return
	}
	if r, err = e.Signature.Marshal(r); chk.E(err) {
		return
	}
	r = separator.Add(r)
	return
}

func (e *E) Unmarshal(data []byte) (r []byte, err error) {

	return
}