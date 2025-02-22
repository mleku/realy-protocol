package signature

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"protocol.realy.lol/pkg/separator"
)

func TestS_Marshal_Unmarshal(t *testing.T) {
	var err error
	for range 10 {
		sig := make([]byte, ed25519.SignatureSize)
		if _, err = rand.Read(sig); chk.E(err) {
			t.Fatal(err)
		}
		var s1 *S
		if s1, err = New(sig); chk.E(err) {
			t.Fatal(err)
		}
		var o []byte
		if o, err = s1.Marshal(nil); chk.E(err) {
			t.Fatal(err)
		}
		o = separator.Add(o)
		s2 := &S{}
		var rem []byte
		if rem, err = s2.Unmarshal(o); chk.E(err) {
			t.Fatal(err)
		}
		if len(rem) > 0 {
			log.I.F("%d %s", len(rem), rem)
		}
		if !bytes.Equal(sig, s2.Signature) {
			t.Fatal("signature did not encode/decode faithfully")
		}

	}
}
