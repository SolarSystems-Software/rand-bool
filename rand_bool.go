package randbool

import (
	"math/rand"
	"sync"
)

// BoolGenerator is a generator capable of generating booleans and bits.
//
// All generators are safe for concurrent usage.
type BoolGenerator struct {
	mu            sync.Mutex
	src           rand.Source
	cache         int64
	remainingBits int
}

// Default is the default BoolGenerator.
//
// Use New to create your own generator.
var Default = New()

// NextBit gets the next random bit.
func (generator *BoolGenerator) NextBit() uint8 {
	generator.mu.Lock()
	defer generator.mu.Unlock()

	// Reset if there's no bits left.
	if generator.remainingBits == 0 {
		generator.cache = generator.src.Int63()
		generator.remainingBits = 63
	}

	// The next bit
	result := uint8(generator.cache & 0x01)

	// Shift over to the right by 1 for the next time this function is called
	generator.cache >>= 1

	// Decrement remaining bits
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
