package bloom

type BloomFilter struct {
	m      uint
	k      uint
	bitmap []bool // bitmap
}

func NewBloomFilter(size uint, k uint) (bf *BloomFilter) {
	bf = &BloomFilter{}
	bf.bitmap = make([]bool, 10)
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
// the list of multiply hashed indexes into hte bitmap
func (bf *BloomFilter) hashindexes(value string) []int {
	return []int{1, 2, 3}
}
