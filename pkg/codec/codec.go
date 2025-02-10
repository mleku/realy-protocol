package codec

// C is an interface for encoding and decoding that allows embedding encoders
// inside other encoders by the use of append for Marshal and slice for
// Unmarshal.
type C interface {
	// Marshal data by appending it to the provided destination, and return the
	// resultant slice.
	Marshal(dst []byte) (r []byte, err error)
	// Unmarshal the next expected data element from the provided slice and return
	// the remainder after the expected separator.
	Unmarshal(data []byte) (r []byte, err error)
}
