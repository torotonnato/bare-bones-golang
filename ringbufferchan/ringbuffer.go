// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

package ringbufferchan

type RingBufferChan[T any] struct {
	InChan, OutChan chan T
}

func NewRingBufferChan[T any](cap int) *RingBufferChan[T] {
	inChan := make(chan T)
	outChan := make(chan T, cap)
	r := RingBufferChan[T]{inChan, outChan}
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
}

func (r *RingBufferChan[T]) Close() {
	close(r.InChan)
	for range r.OutChan {
	} //Wait for worker
	r.InChan, r.OutChan = nil, nil
}
