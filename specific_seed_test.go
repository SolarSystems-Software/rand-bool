package randbool

import (
	"math/rand"
	"testing"
)

func TestGeneration(test *testing.T) {
	src := rand.NewSource(1)
	gen := NewFromSrc(src)

	if gen.NextBool() || !gen.NextBool() || gen.NextBool() {
		test.Error("Bad booleans")
	}
}
