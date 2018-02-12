# sketchbook

Implementing a handful of sketching data structures for the fun of it. :)
See slides for [my talk on sketching data structures](https://speakerdeck.com/kiran/velocity-london-2017-a-tour-of-sketching-data-structures) for context!

### Implemented:

- [Bloom filters](https://en.wikipedia.org/wiki/Bloom_filter) are a probabilistic data structure that attempts to answer membership queries --
ie, is the element in the set. A Bloom filter may return false positives, but never returns a false
negative.

- [HyperLogLogs](https://en.wikipedia.org/wiki/HyperLogLog) count the number of unique elements seen so far in a stream (ie, its cardinality). It uses log(log(cardinality)) space!!
	- HyperLogLog paper: http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf
	- HyperLogLog++ paper: http://research.google.com/pubs/pub40671.html

### coming soon

- [Count Min Sketches](https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch) produces an approximate frequency table of items in a stream. It can overestimate frequency, but will never underestimate. It's very similar to a Bloom Filter in implementation.

- [T-digests](https://github.com/tdunning/t-digest) are an improvement on [Q-digests](https://dl.acm.org/citation.cfm?id=1031495.1031524&coll=DL&dl=GUIDE), which estimate quantiles of a stream with minimal error at the extremes.