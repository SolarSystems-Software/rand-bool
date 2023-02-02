package randbool

import (
	"sync"
	"testing"
)

func TestNextBit(t *testing.T) {
	Default.NextBit()
}

func TestNextBitConcurrent(t *testing.T) {
	var group sync.WaitGroup

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			t.Log("Bit:", Default.NextBit())
			group.Done()
		}()
	}

	group.Wait()
}

func BenchmarkNextBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default.NextBit()
	}
}
