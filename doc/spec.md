# realy protocol event specification

JSON is awful, and space inefficient, and complex to parse due to its intolerance of terminal commas and annoying to work with because of its retarded, multi-standards of string escaping.

Line structured documents are much more readily amenable to human reading and editing, and `\n`/`;`/`:` is more efficient than `","` as an item separator. Data structures can be much more simply expressed in a similar way as how they are in programming languages.

It is one of the guiding principles of the Unix philosophy to keep data in plain text, human readable format wherever possible, forcing the interposition of a parser just for humans to read the data adds extra brittleness to a protocol.

So, this is how realy events look:

```
<type name>\n
<pubkey>\n // encoded in URL-base64
<unix second precision timestamp in decimal ascii>\n
key:value;extra;...\n // zero or more line separated, fields cannot contain a semicolon, end with newline instead of semicolon, key lowercase alphanumeric, first alpha, only key is mandatory, only reserved is `content`
content: // literally this word on one line
<content>\n // any number of further line breaks, last line is signature
<bip-340 schnorr signature encoded in URL-base64>\n
```

The canonical form is exactly this, except for the signature and following linebreak, hashed with Blake2b

The database stored form of this event should make use of an event ID hash to monotonic collision free serial table and an event table.

Event ID hashes will be encoded in URL-base64 where used in tags or mentioned in content with the prefix `event:`. Public keys must be prefixed with `pubkey:` Tag keys should be intelligible words and a specification for their structure should be defined by users of them and shared with other REALY devs.