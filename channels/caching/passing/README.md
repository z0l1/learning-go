# Data passing
## "Cache passing"

Under 'cache' dir and calling it Cache passing because it's for a concurrent in-memory data store in an app of mine.

The motivation is from
[this go101 article](https://go101.org/article/channel.html)
that quotes Rob Pike, saying

> _Don't communicate by sharing data, share data by communicating._

which sounds amazing. The article inspired me to think about a caching solution that works this way.
So after reading it I started thinking
and the first thing that came to was how we used to collect names anywhere.
Bob writes his name on it, passes to Pete and so on.

The idea was pretty simple and so is the implementation.
Probably because it's more like students standing in line to write up their names. So didn't really implement the idea, but I'll call it close enough.

Would love to test this more thoroughly because it seems too good to be true.
Was **suspiciously** easy and works too well with my makeshift mini test.

Will try to `come up` with other concurrent thread-safe stuff
even though there might be a ton of these out there already (including this one) from like 10 years ago `:D`
but where's the fun in that.