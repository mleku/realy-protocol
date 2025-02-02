package types

import (
	"testing"
)

func TestT_Marshal_Unmarshal(t *testing.T) {
	typ := New("note")
	var err error
	var res []byte
	if res, err = typ.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	t2 := new(T)
	var rem []byte
	if rem, err = t2.Unmarshal(res); chk.E(err) {
		t.Fatal(err)
	}
	if len(rem) > 0 {
		log.I.S(rem)
	}
	if !typ.Equal(t2) {
		t.Fatal("types.T did not encode/decode faithfully")
	}
}
