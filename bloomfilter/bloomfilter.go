/*
Package bloom implements Bloom filters.

A Bloom filter is a probabilistic data structure that attempts to answer membership queries --
ie, is the element in the set. A Bloom filter may return false positives, but never returns a false
negative.

A Bloom filter has two tuning parameters: _m_, which is the size of the filter, and _k_, which is the
number of hashes the Bloom filter uses per element. A Bloom filter is backed by a bitset (a bool array
in this implementation.)

The hashing function used in this implementation is murmurhash, a non-cryptographic hashing function.
*/

package bloom

import (
	"hash"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	m      uint
	k      uint
	h      hash.Hash64
	bitmap []bool // bitmap
}

// NewBloomFilter initializes and returns a new Bloom Filter.
func NewBloomFilter(size uint, k uint) (bf *BloomFilter) {
	bf = &BloomFilter{h: murmur3.New64(), k: k, m: size}
	// todo: how do you make a golang bitmap?
	bf.bitmap = make([]bool, size)
	return
}

// Add takes a string value to add, and throws it into the bitmap.
// todo: support any hashable values, not just strings. (or []byte?)
func (bf *BloomFilter) Add(value string) {
	indexes := bf.hashindexes(value)
	// set the bits in the bitmap
	for _, i := range indexes {
		bf.bitmap[i] = true
	}
}

// Test looks for the value in the bloom filter, and returns true
// if it's possible that the value is in the set. It is possible
// for the bloom filter to return true falsely, but a false return
// is always correct.
func (bf *BloomFilter) Test(value string) bool {
	indexes := bf.hashindexes(value)
	// test the bits in the bitmap
	present := true
	for _, hashIndex := range indexes {
		// if any of the indexes isn't set to true, it's not
		// there
		if !(bf.bitmap[hashIndex]) {
			present = false
		}
	}
	return present
}

// hashIndexes takes the value to test for, and returns
// the list of multiply hashed indexes into the bitmap
// Kirsch & Mitzenmacher (2008) show that you can simulate
// i independent hash functions by using g_i(x) = h1(x) + ih2(x)
// instead of using 2 separate hashes, this just splits a 128-bit
// hash in two.
func (bf *BloomFilter) hashindexes(value string) []uint {
	indexes := make([]uint, bf.k, bf.k)

	bf.h.Reset()
	bf.h.Write([]byte(value))
	hashv := bf.h.Sum64()
	// TODO: is this how you split a 64-bit int?
	hashValue1 := uint(hashv)
	hashValue2 := uint(hashv << 32)

	for i := uint(0); i < bf.k; i++ {
		indexes[i] = (hashValue1 + (i * hashValue2)) % bf.m
	}
	return indexes
}
