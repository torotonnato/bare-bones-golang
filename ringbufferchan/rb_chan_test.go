package ringbufferchan

import (
	"testing"
	"time"
)

func testAtCapacity(t *testing.T) {
	cap := 64
	c := NewRingBufferChan[int](cap)
	defer c.Close()

	for it := 0; it < cap; it++ {
		c.WriteChan <- it

	}

	//Wait for the ringbufferchan goroutine to
	//finish. Only needed in testing
	time.Sleep(100 * time.Millisecond)

	for it := 0; it < cap; it++ {
		actual := <-c.ReadChan
		if it != actual {
			t.Errorf("Expected %d, retrieved %d", it, actual)
		}
	}
}

func testAtOverCapacity(t *testing.T) {
	cap := 256
	c := NewRingBufferChan[int](cap)
	defer c.Close()

	for it := 0; it < cap*2; it++ {
		c.WriteChan <- it
	}

	//Wait for the ringbufferchan goroutine to
	//finish. Only needed in testing
	time.Sleep(100 * time.Millisecond)

	for it := 0; it < cap; it++ {
		actual := <-c.ReadChan
		expectation := it + cap
		if actual != expectation {
			t.Errorf("Retrieved: %d, should have been: %d", actual, expectation)
		}
	}
}

func testClose(t *testing.T) {
	c := NewRingBufferChan[int](16)
	c.WriteChan <- 123
	c.Close()

	select {
	case <-c.WriteChan:
		t.Errorf("c.WriteChan is still open")
	case <-c.ReadChan:
		t.Errorf("c.ReadChan is still open")
	default:
	}

	c.Close() //Should be a NOP
}

func TestNewRingBufferChan(t *testing.T) {
	testClose(t)
	testAtCapacity(t)
	testAtOverCapacity(t)
}
