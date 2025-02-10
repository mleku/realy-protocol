package auth

import (
	"protocol.realy.lol/pkg/codec"
	"protocol.realy.lol/pkg/decimal"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/signature"
	"protocol.realy.lol/pkg/url"
)

type Message struct {
	Payload    codec.C
	RequestURL *url.U
	Timestamp  *decimal.T
	PubKey     *pubkey.P
	Signature  *signature.S
}

func SignMessage(msg *Message) (m *Message, err error) {

	return
}

func (m *Message) Marshal(d []byte) (r []byte, err error) {
	r = d
	if r, err = m.Payload.Marshal(d); chk.E(err) {
		return
	}
	if r, err = m.RequestURL.Marshal(d); chk.E(err) {
		return
	}
	if r, err = m.Timestamp.Marshal(d); chk.E(err) {
		return
	}
	if r, err = m.PubKey.Marshal(d); chk.E(err) {
		return
	}
	if r, err = m.Signature.Marshal(d); chk.E(err) {
		return
	}
	return
}

func (m *Message) Unmarshal(d []byte) (r []byte, err error) {

	return
}
