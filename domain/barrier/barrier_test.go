package barrier

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BARRIER_SIZE = 2
)

func TestBarrier(t *testing.T) {
	mutex := sync.RWMutex{}
	barrier := NewBarrier(BARRIER_SIZE)
	count := 0
	barChan := make(chan int)
	endChan := make(chan bool)
	for i := 0; i < BARRIER_SIZE; i++ {
		go func(j int) {
			barChan <- 1
			barrier.Wait()
			mutex.RLock()
			defer mutex.RUnlock()
			if count == 0 {
				endChan <- true
			} else {
				barChan <- -1
			}

		}(i)
	}

	i := 0
	select {
	case x := <-barChan:
		mutex.Lock()
		defer mutex.Unlock()
		count += x
		switch i {
		case 0:
			assert.Equal(t, count, 1)
		case 1:
			assert.Equal(t, count, 2)
		case 2:
			assert.Equal(t, count, 1)
		case 3:
			assert.Equal(t, count, 0)
		}
		i++
	case <-endChan:
		return
	}
}
