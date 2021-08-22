package randbool

import (
	"math/rand"
	"sync"
)

// BoolGenerator can be used to generate random bools and bits.
type BoolGenerator struct {
	sync.Mutex

	src           rand.Source
	cache         int64
	remainingBits int
}

// Default is a generally-available BoolGenerator.
// Use New to create a new BoolGenerator.
var Default = New()

// NextBit gets the next random bit.
func (generator *BoolGenerator) NextBit() uint8 {
	// concurrency safety
	generator.Lock()
	defer generator.Unlock()

	// If there's no bits left, reset
	if generator.remainingBits == 0 {
		generator.cache = generator.src.Int63()
		generator.remainingBits = 63
	}

	// The next bit
	result := uint8(generator.cache & 0x01)

	// Shift over to the right by 1 for the next time this function is called
	generator.cache >>= 1

	// Decrement because we just used one
	generator.remainingBits--

	return result
}

// NextBool gets the next random boolean.
func (generator *BoolGenerator) NextBool() bool {
	return generator.NextBit() == 1
}

// NewFromSrc creates a new BoolGenerator from the given rand.Source.
func NewFromSrc(src rand.Source) *BoolGenerator {
	return &BoolGenerator{
		src:           src,
		cache:         src.Int63(),
		remainingBits: 63,
	}
}

// New creates a new BoolGenerator using a new rand.Source using CryptoRandOrPanic as the seed.
func New() *BoolGenerator {
	return NewFromSrc(rand.NewSource(int64(CryptoRandOrPanic())))
}
