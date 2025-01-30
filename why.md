# why realy?

Since the introduction of the idea of a general "public square" style social network as seen with Facebook and Twitter, the whole world has been overcome by something of a plague of mind control brainwashing cults.

Worse than "Beatlemania" people are being lured into the control of various kinds of "influencers" and adopting in-group words and "challenges" that are more often harmful to the people than actually beneficial like an exercise challenge might be.

Nostr protocol is a super simple event bus architecture, blended with a post office protocol, and due to various reasons related to the recent buyout of Twitter by Elon Musk, who plainly wants to turn it into the Western version of Wechat, it has become plagued with bad subprotocol designs that negate the benefits of self sovereign identity (elliptic curve asymmetric cryptography) and a dominant form of client that is essentially a travesty of Twitter itself.

Realy is being designed with the lessons learned from Nostr and the last 30 years of experience of internet communications protocols to aim to resist this kind of Embrace/Extend/Extinguish protocol that has repeatedly been performed on everything from email, to RSS, to threaded forums and instant messaging, by starting with the distilled essence of how these protocols should work so as to not be so easily vulnerable to being coopted by what is essentially in all but name the same centralised event bus architecture of social networks like Facebook and Twitter.

The main purposes that Realy will target are:

- synchronous instant messaging protocols with IRC style nickserv and chanserv permissions and persistence, built from the ground up to take advantage of the cryptographic identities created by BIP-340 signatures, with an intuitive threaded structure that allows users to peruse a larger discussion without the problem of threads of discussion breaking the top level structure
- structured document repositories primarily for text media, as a basis for collaborative documentation and literature collections, and software source code (breaking out of the filesystem tree structure to permit much more flexible ways of organising code)
- persistent threaded discussion forums for longer form messages than the typical single sentence/paragraph of instant messaging
- simple cross-relay data query protocol that enables minimising the data cost of traffic to clients
- push style notification systems that can be programmed by the users' clients to respond to any kind of event breadcast to a relay

A key concept in the R.E.A.L.Y. architecture is that of relays being a heteregenous group of data repositories and relaying systems that are built specific to purpose, such as a chat relay, which does not store any messages but merely bounces messages around ot subscribers, a document repository, which provides read access to data with full text search capability, that can ne specialised for a singular data format (eg markdown, eg mediawiki, eg code), a threaded, moderated forum, and others.

A second key concept in R.E.A.L.Y. is the integration of Lightning Network payments - again mostly copying what is done with Nostr but enabling both per-use, micro-accounts and long term subscription styles of access, and the promotion of a notion of user-pays - where all data writing must be charged for, and most reading must be paid for. Lightning is perfect for this because it can currently cope with enormous volumes of payments with mere seconds of delay for settlement and a granularity of denomination that lends itself to  the very low cost of delivering a one-time service, or maintaining a micro-account.
