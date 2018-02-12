package hyperloglog

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountLeadingZeros(t *testing.T) {
	h, _ := Initialize(8)
	assert.Equal(t, uint8(0), h.countLeadingZeros(0x91<<56))
	assert.Equal(t, uint8(0), h.countLeadingZeros(0xab<<56))
	assert.Equal(t, uint8(1), h.countLeadingZeros(0x48<<56))
	assert.Equal(t, uint8(2), h.countLeadingZeros(0x28<<56))
	assert.Equal(t, uint8(3), h.countLeadingZeros(0x18<<56))
	assert.Equal(t, uint8(4), h.countLeadingZeros(0x09<<56))
	assert.Equal(t, uint8(5), h.countLeadingZeros(0x04<<56))
}

func TestSplitWord(t *testing.T) {
	h, _ := Initialize(16)

	i, remainder := h.splitWord(0x1111222233334444)
	assert.Equal(t, uint64(0x1111), i)
	assert.Equal(t, uint64(0x2222333344440000), remainder)

	i, remainder = h.splitWord(0x0123456789abcdef)
	assert.Equal(t, uint64(0x0123), i)
	assert.Equal(t, uint64(0x456789abcdef0000), remainder)
}

func TestHarmonicMean(t *testing.T) {
	nums := []uint8{1, 4, 4}
	assert.Equal(t, float64(2), harmonicMean(nums))
}

func randStr(n int) string {
	i := rand.Uint32()
	return fmt.Sprintf("a%s %s", i, n)
}

func TestHll(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	hll, _ := Initialize(16)
	n := 300000
	for i := 0; i < n; i++ {
		val := rand.Int()
		hll.Add(randStr(val))
		hll.Add(randStr(val))
	}
	cardinalityEstimate := hll.Cardinality()

	assert.InEpsilon(t, n, cardinalityEstimate, 0.03)
}
