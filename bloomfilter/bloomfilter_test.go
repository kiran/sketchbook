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
	bf := bloom.NewBloomFilter(10, 2)
	bf.Add("hello")
	assert.True(t, bf.Test("hello"), "Expected the element to be in the bloom filter")
}
