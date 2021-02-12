package randbool

import "testing"

func TestDefaultImplementation(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log("Bit:", Default.NextBit())
	}
}
