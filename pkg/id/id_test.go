package id

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"protocol.realy.lol/pkg/separator"
)

func TestT_Marshal_Unmarshal(t *testing.T) {
	var err error
	for range 10 {
		pk := make([]byte, ed25519.PublicKeySize)
		if _, err = rand.Read(pk); chk.E(err) {
			t.Fatal(err)
		}
		var p *T
		if p, err = New(pk); chk.E(err) {
			t.Fatal(err)
		}
		var o []byte
		if o, err = p.Marshal(nil); chk.E(err) {
			t.Fatal(err)
		}
		o = separator.Add(o)
		p2 := &T{}
		var rem []byte
		if rem, err = p2.Unmarshal(o); chk.E(err) {
			t.Fatal(err)
		}
		if len(rem) > 0 {
			log.I.F("%d %s", len(rem), rem)
		}
		if !bytes.Equal(pk, p2.b) {
			t.Fatal("public key did not encode/decode faithfully")
		}
	}
}