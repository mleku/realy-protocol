package content

import (
	"bytes"
	"crypto/rand"
	mrand "math/rand"
	"testing"

	"protocol.realy.lol/pkg/separator"
)

func TestC_Marshal_Unmarshal(t *testing.T) {
	c := make([]byte, mrand.Intn(100)+25)
	_, err := rand.Read(c)
	if err != nil {
		t.Fatal(err)
	}
	c1 := new(C)
	c1.Content = c
	var res []byte
	if res, err = c1.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	res = separator.Add(res)
	c2 := new(C)
	var rem []byte
	if rem, err = c2.Unmarshal(res); chk.E(err) {
		t.Fatal(err)
	}
	if !bytes.Equal(c1.Content, c2.Content) {
		log.I.S(c1.Content, c2.Content)
		t.Fatal("content not equal")
	}
	if len(rem) > 0 {
		log.I.S(rem)
		t.Fatalf("unexpected remaining bytes: '%0x'", rem)
	}
}