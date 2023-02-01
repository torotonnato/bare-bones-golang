package ringbufferchan

import (
	"testing"
	"time"
)

func testSingleItem(t *testing.T) {
	c := NewRingBufferChan[int](64)
	defer c.Close()
	item := 123
	c.InChan <- item
	if <-c.OutChan != item {
		t.Errorf("Single item test failed")
	}
}

func testCapacityOverflow(t *testing.T) {
	cap := 256
	c := NewRingBufferChan[int](cap)
	defer c.Close()
	for it := 0; it < cap*2; it++ {
		c.InChan <- it
	}
	//Wait for the ringbufferchan goroutine to
	//finish. Only needed in testing
	time.Sleep(100 * time.Millisecond)
	for it := 0; it < cap; it++ {
		actual := <-c.OutChan
		expectation := it + cap
		if actual != expectation {
			t.Errorf("Retrieved: %d, should have been: %d", actual, expectation)
		}
	}
}

func testClose(t *testing.T) {
	c := NewRingBufferChan[int](16)
	c.InChan <- 123
	c.Close()
	select {
	case <-c.InChan:
		t.Errorf("c.InChan is still open")
	case <-c.OutChan:
		t.Errorf("c.OutChan is still open")
	default:
	}
}

func TestNewRingBufferChan(t *testing.T) {
	testClose(t)
	testSingleItem(t)
	testCapacityOverflow(t)
}
