package counter_test

import (
	"gkgk-go-iv/counter"
	"sync"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCounter(t *testing.T) {
	var wg sync.WaitGroup
	counter := counter.NewCounter()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	expected := int64(100)
	actual := counter.GetCount()
	assert.Equal(t, actual, expected)
}

func TestCounterIncrDecr(t *testing.T) {
	var wg sync.WaitGroup
	counter := counter.NewCounter()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			if j%2 == 0 {
				counter.Increment()
			} else {
				counter.Decrement()
			}
		}(i)
	}
	wg.Wait()

	expected := int64(0)
	actual := counter.GetCount()
	assert.Equal(t, actual, expected)
}
