package main

// given a stream of integers, and a set amount of memory
// allocate a map of m bytes, which are buckets of first log(m) bytes
// hash the numbers coming in, and find the bucket to put them in
// calculate the stream of 0s at the beginning of the next few bytes
// and record the maximum number seen into the correct bucket

// HyperLogLog estimates cardinality
type HyperLogLog struct {
	// has some stuff I guess
	buckets []byte // number of bytes to allocate
	m uint32
	p uint8
}

// ProcessNumber adds another number to the HyperLogLog structure
func (hll *HyperLogLog) ProcessNumber(number int) {
	// do some stuff
	return
}

// Cardinality emits the current estimated cardinality
func (hll *HyperLogLog) Cardinality() (cardinality int) {
	return 0
}

// Initialize takes a number of bytes to initialize with,
// and returns the struct
func Initialize(m int) *HyperLogLog {
	hll := &HyperLogLog{}
	hll.buckets = make([]byte, m)
	return hll
}

// EstimatedError returns the approximate precision
func (hll *HyperLogLog) EstimatedError() (est float64) {
	return 0.0
}

func main() {
	// do nothing
}