package hyperloglog

import (
	"errors"
	"hash/fnv"
	"math"
)

// given a stream of integers, and a set amount of memory
// allocate a map of m bytes, which are buckets of first log(m) bytes
// hash the numbers coming in, and find the bucket to put them in
// calculate the stream of 0s at the beginning of the next few bytes
// and record the maximum number seen into the correct bucket

// HyperLogLog estimates cardinality
type HyperLogLog struct {
	buckets     []uint8 // byte array of counts
	bucketCount uint64  // number of buckets
	precision   uint8   // precision (number of bytes to count)
	debug       bool
}

// Initialize takes a number of bytes to initialize with,
// and returns the struct
func Initialize(precision uint8) (*HyperLogLog, error) {
	hll := &HyperLogLog{}
	if precision > 16 || precision < 4 {
		return nil, errors.New("precision must be between 4 and 16")
	}
	hll.precision = precision
	hll.bucketCount = 1 << precision
	hll.buckets = make([]uint8, hll.bucketCount)
	return hll, nil
}

// Add adds another number to the HyperLogLog structure
func (hll *HyperLogLog) Add(value string) {
	// hash the number
	h := fnv.New64a()
	h.Write([]byte(value))
	hashedNumber := h.Sum64()
	// a. figure out which bucket to put it into
	bucket, remainder := hll.splitWord(hashedNumber)
	// b. count the longest run of leading zeroes in the second half
	contiguousZeros := hll.countLeadingZeros(remainder)

	if contiguousZeros > hll.buckets[bucket] {
		hll.buckets[bucket] = contiguousZeros
	}
	return
}

// splitWord, given a hashedNumber, extracts the leading bits to determine
// which bucket the number should contribute to
func (hll *HyperLogLog) splitWord(hashedNumber uint64) (uint64, uint64) {
	// given a hashed number, figure out which bucket it goes into

	// shift the hashed number over until precision bits are left
	// then mask out the rest of the number
	i := (hashedNumber >> (64 - hll.precision)) & (hll.bucketCount - 1)

	// the remainder --
	remainder := (hashedNumber << hll.precision)

	return i, remainder
}

func (hll *HyperLogLog) countLeadingZeros(remainder uint64) uint8 {
	// given a hashed number, figure out how many leading zeros it has
	var mask uint64
	mask = 1 << 63
	clz := uint8(0)
	// limit = 1 << (hll.p - 1)
	for (remainder & mask) == 0 {
		mask = mask >> 1
		clz++
	}
	return clz
}

func harmonicMean(nums []uint8) float64 {
	var invertedSum float64
	for _, num := range nums {
		invertedSum += 1.0 / float64(num)
	}
	return float64(len(nums)) / invertedSum
}

// Cardinality emits the current estimated cardinality
func (hll *HyperLogLog) Cardinality() (cardinalityEstimate float64) {
	// if there are empty buckets, use linear counting
	// cardinalityEstimate = harmonicMean(hll.buckets)

	invertedSum := 0.0
	for _, zeros := range hll.buckets {
		invertedSum += 1.0 / math.Exp2(float64(zeros))
	}

	est := 0.7213 / (1 + 1.079/float64(hll.bucketCount))

	cardinalityEstimate = est * float64(hll.bucketCount) * float64(hll.bucketCount) / invertedSum

	// else do some low-range/high-range bias correction

	// count up all the HLL zeros estimates
	// take the harmonic mean of everything; renormalize
	return cardinalityEstimate
}

// EstimatedError returns the approximate precision
func (hll *HyperLogLog) EstimatedError() (est float64) {
	// the standard error for an HLL with n registers is less than 1.04/sqrt(m).
	return 1.04 / (math.Sqrt(float64(hll.bucketCount)))
}
