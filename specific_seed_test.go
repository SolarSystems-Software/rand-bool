package randbool

import (
	"math/rand"
	"testing"
)

func TestGeneration(test *testing.T) {
	src := rand.NewSource(1)
	gen := NewFromSrc(src)

	if gen.NextBoolean() || !gen.NextBoolean() || gen.NextBoolean() {
		test.Error("Bad booleans")
	}
}
