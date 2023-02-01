// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

package ringbufferchan

import "sync"

type RingBufferChan[T any] struct {
	WriteChan, ReadChan chan T
	sync.WaitGroup
}

func NewRingBufferChan[T any](cap int) *RingBufferChan[T] {
	r := RingBufferChan[T]{
		WriteChan: make(chan T),
		ReadChan:  make(chan T, cap),
	}
	r.Add(1)
	go r.run()
	return &r
}

func (r *RingBufferChan[T]) run() {
	for v := range r.WriteChan {
		select {
		case r.ReadChan <- v:
		default:
			<-r.ReadChan
			r.ReadChan <- v
		}
	}
	close(r.ReadChan)
	r.Done()
}

func (r *RingBufferChan[T]) Close() {
	if r.WriteChan != nil {
		close(r.WriteChan)
		r.Wait()
		r.WriteChan, r.ReadChan = nil, nil
	}
}
