package RandBool

import (
	"Crimson/RandBool/randbool"
	"math/rand"
	"testing"
)

func TestGeneration(test *testing.T) {
	src := rand.NewSource(1)
	gen := randbool.NewFromSrc(src)

	if randbool.NextBoolean(gen) || !randbool.NextBoolean(gen) || randbool.NextBoolean(gen) {
		test.Error("Bad booleans")
	}
}
