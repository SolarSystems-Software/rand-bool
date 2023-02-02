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
	cache         int64
	remainingBits int
}

// Default is the default BoolGenerator.
var Default = BoolGenerator{}

// NextBit gets the next random bit.
func (generator *BoolGenerator) NextBit() uint8 {
	generator.mu.Lock()
	defer generator.mu.Unlock()

	// Reset if there's no bits left.
	if generator.remainingBits == 0 {
		generator.cache = rand.Int63()
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
