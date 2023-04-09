package barrier

import "sync"

type Barrier struct {
	size, count int
	mutex       *sync.Mutex
	cond        *sync.Cond
}

func NewBarrier(size int) *Barrier {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	return &Barrier{
		size, 0, mutex, cond,
	}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.count++
	if b.count == b.size {
		b.count = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
}
