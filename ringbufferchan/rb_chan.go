// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

package ringbufferchan

import "sync"

type RingBufferChan[T any] struct {
	InChan, OutChan chan T
	sync.WaitGroup
}

func NewRingBufferChan[T any](cap int) *RingBufferChan[T] {
	r := RingBufferChan[T]{
		InChan:  make(chan T),
		OutChan: make(chan T, cap),
	}
	r.Add(1)
	go r.run()
	return &r
}

func (r *RingBufferChan[T]) run() {
	for v := range r.InChan {
		select {
		case r.OutChan <- v:
		default:
			<-r.OutChan
			r.OutChan <- v
		}
	}
	close(r.OutChan)
	r.Done()
}

func (r *RingBufferChan[T]) Close() {
	if r.InChan != nil {
		close(r.InChan)
		r.Wait()
		r.InChan, r.OutChan = nil, nil
	}
}
