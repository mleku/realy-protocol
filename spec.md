# realy protocol event specification

JSON is awful, and space inefficient, and complex to parse due to its intolerance of terminal commas and annoying to work with because of its retarded, multi-standards of string escaping.

Line structured documents are much more readily amenable to human reading and editing, and `\n` is more efficient than `","` as an item separator. Data structures can be much more simply expressed in a similar way as how they are in programming languages.

It is one of the guiding principles of the Unix philosophy to keep data in plain text, human readable format wherever possible, forcing the interposition of a parser just for humans to read the data adds extra brittleness to a protocol.

So, this is how realy events look:

```
<type name>\n
<pubkey>\n
<unix second precision timestamp in decimal ascii>\n
key;value;extra;...\n // zero or more line separated, fields cannot contain a semicolon, end with newline instead of semicolon
content // literally this word on one line
<content>\n
<bip-340 schnorr signature encoded in base64> 
```
