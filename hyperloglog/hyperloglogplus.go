package main

import (
	"encoding/binary"
	"fmt"
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
	buckets     []byte // byte array of buckets
	bucketCount uint32 // number of buckets
	precision   uint8  // precision (number of bytes to count)
}

// ProcessNumber adds another number to the HyperLogLog structure
func (hll *HyperLogLog) ProcessElement(value string) {
	// hash the number
	h := fnv.New64a()
	h.Write([]byte(value))
	hashedNumber := h.Sum64()
	fmt.Printf("%+v\n", hashedNumber)
	// slice the number --
	hashedBytes := convertToBytes(hashedNumber)
	// a. figure out which bucket to put it into
	bucket := hll.getBucket(hashedBytes)
	// b. count the longest run of leading zeroes in the second half
	contiguousZeros := hll.countZeros(hashedBytes)

	if contiguousZeros > hll.buckets[bucket] {
		hll.buckets[bucket] = contiguousZeros
	}
	return
}

func convertToBytes(hashedNumber uint64) []byte {
	bs := make([]byte, 64)
	binary.LittleEndian.PutUint64(bs, hashedNumber)
	fmt.Println(bs)
	return bs
}

func (hll *HyperLogLog) getBucket(hashedBytes []byte) int {
	// given a hashed number, figure out which bucket it goes into

	return 0
}

func (hll *HyperLogLog) countZeros(hashedBytes []byte) uint8 {
	// given a hashed number, figure out how many leading zeros it has
	return 0
}

// Cardinality emits the current estimated cardinality
func (hll *HyperLogLog) Cardinality() (cardinality int) {
	// if there are empty buckets, use linear counting

	// else do some low-range/high-range bias correction

	// count up all the HLL zeros estimates
	// take the harmonic mean of everything; renormalize
	return 0
}

// Initialize takes a number of bytes to initialize with,
// and returns the struct
func Initialize(precision uint8) *HyperLogLog {
	hll := &HyperLogLog{}
	// todo: is this cast safe?
	hll.precision = precision
	hll.bucketCount = 1 << precision
	hll.buckets = make([]byte, hll.bucketCount)
	return hll
}

// EstimatedError returns the approximate precision
func (hll *HyperLogLog) EstimatedError() (est float64) {
	// the standard error for an HLL with n registers is less than 1.04/sqrt(m).
	return 1.04 / (math.Sqrt(float64(hll.bucketCount)))
}

func main() {
	h := Initialize(3)
	h.ProcessElement("kittens")
	h.ProcessElement("elephants")
}
