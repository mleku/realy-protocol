package event

import (
	"encoding/binary"
	mrand "math/rand"
	"math/rand/v2"
	"testing"
	"time"

	"lukechampine.com/frand"
	"realy.lol/signer"

	"protocol.realy.lol/pkg/content"
	"protocol.realy.lol/pkg/decimal"
	"protocol.realy.lol/pkg/id"
	"protocol.realy.lol/pkg/pubkey"
	"protocol.realy.lol/pkg/tag"
	"protocol.realy.lol/pkg/tags"
	"protocol.realy.lol/pkg/types"

	"realy.lol/p256k"
)

const seed = 0

func GenerateFake32Bytes(rng *rand.Rand) (fake []byte) {
	fake = make([]byte, 32)
	for i := range 4 {
		n := rng.Uint64()
		binary.LittleEndian.PutUint64(fake[i*8:i*8+8], n)
	}
	return
}

var Hashtags, _ = tag.New(
	"halsey",
	"$DIAM",
	"Trevor Lawrence",
	"#AEWCEO",
	"Reuters",
	"Linda McMahon",
	"Bolton",
	"Raining in Houston",
	"#SwiftDay",
	"Munich",
	"NATO",
	"#thursdayvibes",
	"Good Thursday",
	"$SEA",
	"#AEWGrandSlam",
	"Brian Steele",
	"#GalentinesDay",
	"Bregman",
	"Afghan",
	"The Accountant 2",
	"Happy Friday Eve",
	"TLaw",
	"Red Sox",
	"Large Scale Social Deception",
	"2024 BMW",
	"Onew",
	"Secretary of Education",
	"$HIMS",
	"Core PPI",
	"Avowed",
	"Kemp",
	"Angel's Venture",
	"YouTube TV",
	"Bri Bri",
	"Teslas",
	"Thirsty Thursday",
	"matz",
	"Jack the Ripper",
	"Paramount",
	"Megan Boswell",
	"Zeldin",
	"Zelensky",
	"Censure",
	"Sheldon Whitehouse",
	"Arenado",
	"Parasite Class",
	"Kennedy Center",
	"I Love Jesus",
	"James Cook",
)

func GenerateContent(rng *rand.Rand, l int) (c *content.C) {
	c = &content.C{}

	return
}

const lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 

Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func GenerateTags(rng *rand.Rand, n int) (t *tags.T, err error) {
	nE, nP, nH := rng.IntN(n)+1, rng.IntN(n)+1, rng.IntN(n)+1
	var tt []*tag.T
	k := tag.List.GetElementBytes(tag.KeyEvent)
	for range nE {
		var tg *tag.T
		v := GenerateFake32Bytes(rng)
		var e *id.T
		if e, err = id.New(v); chk.E(err) {
			return
		}
		var b []byte
		if b, err = e.Marshal(b); chk.E(err) {
			return
		}
		if tg, err = tag.New(k, b, []byte("root")); chk.E(err) {
			return
		}
		tt = append(tt, tg)
	}
	k = tag.List.GetElementBytes(tag.KeyPubkey)
	for range nP {
		var tg *tag.T
		v := GenerateFake32Bytes(rng)
		var p *pubkey.P
		if p, err = pubkey.New(v); chk.E(err) {
			return
		}
		var b []byte
		if b, err = p.Marshal(b); chk.E(err) {
			return
		}
		if tg, err = tag.New(k, b); chk.E(err) {
			return
		}
		tt = append(tt, tg)
	}
	k = tag.List.GetElementBytes(tag.KeyHashtag)
	for range nH {
		var tg *tag.T
		v := Hashtags.GetElementBytes(rng.IntN(Hashtags.Len() - 1))
		// v = bytes.ReplaceAll(v, []byte{';'}, []byte{'_'})
		// v = bytes.ReplaceAll(v, []byte{':'}, []byte{'-'})
		// log.I.S(v)
		if tg, err = tag.New(k, v); chk.E(err) {
			return
		}
		tt = append(tt, tg)
	}
	t = tags.New(tt...)
	return
}

func GenerateEvent(sign signer.I) (ev *E, err error) {
	s2 := rand.NewPCG(seed, seed)
	rng := rand.New(s2)
	sign = new(p256k.Signer)
	if err = sign.Generate(); chk.E(err) {
		return
	}
	var pk *pubkey.P
	if pk, err = pubkey.New(sign.Pub()); chk.E(err) {
		return
	}
	var t *tags.T
	if t, err = GenerateTags(rng, 3+1); chk.E(err) {
		return
	}
	cont := make([]byte, mrand.Intn(100)+25)
	_, err = frand.Read(cont)

	ev = &E{
		Type:      types.New("note/adoc"),
		Pubkey:    pk,
		Timestamp: decimal.New(time.Now().Unix()),
		Tags:      t,
		Content:   &content.C{Content: []byte(lorem)},
	}
	if err = ev.Sign(sign); chk.E(err) {
		return
	}
	return
}

func TestE_Marshal_Unmarshal(t *testing.T) {
	var ev *E
	var err error
	var b1, b2 []byte
	sign := &p256k.Signer{}
	if err = sign.Generate(); chk.E(err) {
		t.Fatal(err)
	}
	for range 10 {
		if ev, err = GenerateEvent(sign); chk.E(err) {
			t.Fatal(err)
		}
		if b1, err = ev.Marshal(b1); chk.E(err) {
			t.Fatal(err)
		}
		log.I.F("\n```\n%s```\n", b1)
		// log.I.S(ev)
		b1 = b1[:0]
		_ = b2
	}
}
