package content

import (
	"bytes"
	"crypto/rand"
	mrand "math/rand"
	"testing"
)

func TestC_Marshal_Unmarshal(t *testing.T) {
	c := make([]byte, mrand.Intn(100)+25)
	_, err := rand.Read(c)
	if err != nil {
		t.Fatal(err)
	}
	log.I.S(c)
	c1 := new(C)
	c1.Content = c
	var res []byte
	if res, err = c1.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	// append a fake zero length signature
	res = append(res, '\n')
	log.I.S(res)
	c2 := new(C)
	var rem []byte
	if rem, err = c2.Unmarshal(res); chk.E(err) {
		t.Fatal(err)
	}
	if !bytes.Equal(c1.Content, c2.Content) {
		log.I.S(c1, c2)
		t.Fatal("content not equal")
	}
	if !bytes.Equal(rem, []byte{'\n'}) {
		log.I.S(rem)
		t.Fatalf("remainder not found")
	}
}
