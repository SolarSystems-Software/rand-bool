package randbool

import "testing"

func BenchmarkBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.NextBit()
	}
}
