package randbool

import (
	"math/rand"
	"time"
)

// Represents a generator. Used to generate random bits.
type BoolGenerator struct {
	src rand.Source
	cache int64
	remainingBits int
}

// Gets the next bit in the specified BoolGenerator.
func NextBit(generator *BoolGenerator) uint8 {
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

// Gets the next boolean in the specified BoolGenerator.
func NextBoolean(generator *BoolGenerator) bool {
	return NextBit(generator) == 1
}

// Generates a new BoolGenerator using the given rand.Source.
func NewFromSrc(src rand.Source) *BoolGenerator {
	return &BoolGenerator{
		src: src,
		cache: src.Int63(),
		remainingBits: 63,
	}
}

// Generates a new BoolGenerator using a rand.Source from the current UNIX time.
func New() *BoolGenerator {
	return NewFromSrc(rand.NewSource(time.Now().UnixNano()))
}
