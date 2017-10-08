package bloom

import (
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	m      uint
	k      uint
	h1     hash.Hash64
	h2     hash.Hash64
	bitmap []bool // bitmap
}

func NewBloomFilter(size uint, k uint) (bf *BloomFilter) {
	bf = &BloomFilter{h1: fnv.New64(), h2: murmur3.New64(), k: k, m: size}
	bf.bitmap = make([]bool, size)
	return
}

// todo: how do you make a golang bitmap?
func (bf *BloomFilter) Add(value string) {
	indexes := bf.hashindexes(value)
	// set the bits in the bitmap
	for _, i := range indexes {
		bf.bitmap[i] = true
	}
}

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
func (bf *BloomFilter) hashindexes(value string) []uint {
	indexes := make([]uint, bf.k, bf.k)

	bf.h1.Reset()
	bf.h1.Write([]byte(value))
	hashValue1 := uint(bf.h1.Sum64())

	bf.h2.Reset()
	bf.h2.Write([]byte(value))
	hashValue2 := uint(bf.h2.Sum64())

	for i := uint(0); i < bf.k; i++ {
		indexes[i] = (hashValue1 + (i * hashValue2)) % bf.m
	}
	return indexes
}
