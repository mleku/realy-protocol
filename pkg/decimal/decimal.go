package decimal

import (
	_ "embed"
	"time"

	"golang.org/x/exp/constraints"
)

// run this to regenerate (pointlessly) the base 10 array of 0 to 9999
//go:generate go run ./gen/.

//go:embed base10k.txt
var base10k []byte

const base = 10000

type T struct{ N uint64 }

func New[V constraints.Integer](n V) *T { return &T{uint64(n)} }

func Now() *T { return New(time.Now().Unix()) }

func (n *T) Uint64() uint64 { return n.N }
func (n *T) Int64() int64   { return int64(n.N) }
func (n *T) Uint16() uint16 { return uint16(n.N) }

var powers = []*T{
	{base / base},
	{base},
	{base * base},
	{base * base * base},
	{base * base * base * base},
}

const zero = '0'
const nine = '9'

func (n *T) Marshal(d []byte) (r []byte, err error) {
	if n == nil {
		err = errorf.E("cannot marshal nil timestamp")
		return
	}
	nn := n.N
	r = d
	if n.N == 0 {
		r = append(r, '0')
		return
	}
	var i int
	var trimmed bool
	k := len(powers)
	for k > 0 {
		k--
		q := n.N / powers[k].N
		if !trimmed && q == 0 {
			continue
		}
		offset := q * 4
		bb := base10k[offset : offset+4]
		if !trimmed {
			for i = range bb {
				if bb[i] != '0' {
					bb = bb[i:]
					trimmed = true
					break
				}
			}
		}
		r = append(r, bb...)
		n.N = n.N - q*powers[k].N
	}
	// r = append(r, '\n')
	n.N = nn
	return
}

// Unmarshal reads a string, which must be a positive integer no larger than math.MaxUint64,
// skipping any non-numeric content before it.
//
// Note that leading zeros are not considered valid, but basically no such thing as machine
// generated JSON integers with leading zeroes. Until this is disproven, this is the fastest way
// to read a positive json integer, and a leading zero is decoded as a zero, and the remainder
// returned.
func (n *T) Unmarshal(d []byte) (r []byte, err error) {
	if len(d) < 1 {
		err = errorf.E("zero length number")
		return
	}
	var sLen int
	if d[0] == zero {
		r = d[1:]
		n.N = 0
		return
	}
	// count the digits
	for ; sLen < len(d) && d[sLen] >= zero && d[sLen] <= nine && d[sLen] != ','; sLen++ {
	}
	if sLen == 0 {
		err = errorf.E("zero length number")
		return
	}
	if sLen > 20 {
		err = errorf.E("too big number for uint64")
		return
	}
	// the length of the string found
	r = d[sLen:]
	d = d[:sLen]
	for _, ch := range d {
		ch -= zero
		n.N = n.N*10 + uint64(ch)
	}
	return
}