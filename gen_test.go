package RandBool

import (
	"Crimson/RandBool/randbool"
	"math/rand"
	"testing"
)

func TestGeneration(test *testing.T) {
	src := rand.NewSource(1)
	gen := randbool.NewFromSrc(src)

	if gen.NextBoolean() || !gen.NextBoolean() || gen.NextBoolean() {
		test.Error("Bad booleans")
	}
}
