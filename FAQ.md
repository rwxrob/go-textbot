Why force use of jsoncache for responder context?
---------------------------------------------------

There are two approaches. One is to stay really generic and make the
context a `map[string]interface{}`. The other is to lock down the type
to something very specific and manageable. The first would require
additional management libraries that took `map[string]interface{}` as a
parameter. The second allows the addition of methods to the specific
type itself.

The second option, while more constricting, provides a solid avenue for
expansion and improvement without too much additional complexity. The
community can lobby for additions to the `jsoncache` module and they can
be added in a controlled way. The alternative would allow (and
indirectly promote) the creation of responder packs that do not deal
with `map[string]interface{}` in a concurrency safe way.

The initial team discussed this at length and agreed the second option
is more sustainable as well as easier to understand even for beginners.

What if my UUID is not unique or someone spoofs mine?
-----------------------------------------------------

The current plan is for a searchable online database rather than any
sort of registry. This is identical to the `godoc.org` approach with the
addition of some way for responder authors to submit new additions for
indexing, sort of like Google bot.

The first time a UUID is used for a specific package it cannot be used
again, for anything other than that package at that exact location. This
allows the original responder author to update and modify the responder
with full control and others to be sure that UUID is never used by
anyone else. The only way this system can be abused is if access to the
package itself is compromised (which is already an acceptable risk of
using imports in the first place).

What is stopping another responder from altering the cache of another?
----------------------------------------------------------------------

Nothing. In fact, this is allowed and encouraged as a means of
communication between responders. It is up to responder creators to
establish their own context APIs for communication.

You can think of the responder context as a sort of shared memory model
as is used by operating systems themselves for communication between
processes.
