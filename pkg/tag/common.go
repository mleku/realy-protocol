package tag

const (
	KeyEvent = iota
	KeyPubkey
	KeyHashtag
)

var List, _ = New(
	// event is a reference to an event, the value is an Event Id
	"event",
	// pubkey is a reference to a public key, the value is a pubkey.P
	"pubkey",
	// hashtag is a string that can be searched by a hashtag filter tag
	"hashtag",
	// ... can many more things be in here for purposes
)
