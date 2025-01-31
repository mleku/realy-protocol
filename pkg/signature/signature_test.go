package signature

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"
)

func TestS_Marshal_Unmarshal(t *testing.T) {
	var err error
	for range 10 {
		sig := make([]byte, ed25519.SignatureSize)
		if _, err = rand.Read(sig); chk.E(err) {
			t.Fatal(err)
		}
		log.I.S(sig)
		var s *S
		if s, err = New(sig); chk.E(err) {
			t.Fatal(err)
		}
		var o []byte
		if o, err = s.Marshal(nil); chk.E(err) {
			t.Fatal(err)
		}
		log.I.F("%d %s", len(o), o)
		p2 := &S{}
		var rem []byte
		if rem, err = p2.Unmarshal(o); chk.E(err) {
			t.Fatal(err)
		}
		if len(rem) > 0 {
			log.I.F("%d %s", len(rem), rem)
		}
		log.I.S(p2.Signature)
		if !bytes.Equal(sig, p2.Signature) {
			t.Fatal("signature did not encode/decode faithfully")
		}

	}
}
