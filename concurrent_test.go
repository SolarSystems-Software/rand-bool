package randbool

import (
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
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
