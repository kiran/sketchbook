package bloom_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	bloom "github.com/kiran/sketchbook/bloomfilter"
)

func TestInsertion(t *testing.T) {
	bf := bloom.NewBloomFilter(10, 2)
	bf.Add("hello")
}

func TestInclusion(t *testing.T) {
	bf := bloom.NewBloomFilter(100, 2)
	bf.Add("hello")
	assert.True(t, bf.Test("hello"), "Expected the element to be in the bloom filter")

	assert.False(t, bf.Test("henlo"), "Expected the element to _not_ be in the bloom filter")
}

func BenchmarkAddition(b *testing.B) {
	bf := bloom.NewBloomFilter(100, 2)
	for i := 0; i < b.N; i++ {
		bf.Add(string(i))
	}
}
