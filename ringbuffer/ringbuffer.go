// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

package ringbufferchan

type RingBufferChan[T any] struct {
	inChan  <-chan T
	outChan chan T
}

func NewRingBufferChan[T any](inChan <-chan T, outChan chan T) *RingBufferChan[T] {
	return &RingBufferChan[T]{inChan, outChan}
}

func (r *RingBufferChan[T]) Run() {
	for v := range r.inChan {
		select {
		case r.outChan <- v:
		default:
			<-r.outChan
			r.outChan <- v
		}
	}
	close(r.outChan)
}
