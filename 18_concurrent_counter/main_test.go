package main

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AtomicCounter(t *testing.T) {
	ac := atomicCounter{
		counter: &atomic.Uint64{},
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			ac.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, uint64(10000), ac.Value())
}

func Test_MutexCounter(t *testing.T) {
	mc := mutexCounter{
		mu:      &sync.Mutex{},
		counter: 0,
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			mc.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, uint64(10000), mc.Value())
}

func Test_ChannelCounter(t *testing.T) {
	ctx, cf := context.WithCancel(context.Background())
	defer cf()

	cc := newChannelCounter(ctx)

	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			cc.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, uint64(10000), cc.Value())
}

func Benchmark_Counter(b *testing.B) {
	b.Run("mutex_counter", func(b *testing.B) {
		mc := mutexCounter{
			mu:      &sync.Mutex{},
			counter: 0,
		}

		for i := 0; i < b.N; i++ {

			wg := &sync.WaitGroup{}
			for i := 0; i < 10000; i++ {
				wg.Add(1)
				go func() {
					mc.Inc()
					wg.Done()
				}()
			}

			wg.Wait()
		}
	})

	b.Run("atomic_counter", func(b *testing.B) {
		ac := atomicCounter{
			counter: &atomic.Uint64{},
		}

		for i := 0; i < b.N; i++ {

			wg := &sync.WaitGroup{}
			for i := 0; i < 10000; i++ {
				wg.Add(1)
				go func() {
					ac.Inc()
					wg.Done()
				}()
			}

			wg.Wait()
		}
	})

	b.Run("channel_counter", func(b *testing.B) {
		ctx, cf := context.WithCancel(context.Background())
		defer cf()

		cc := newChannelCounter(ctx)

		for i := 0; i < b.N; i++ {

			wg := &sync.WaitGroup{}
			for i := 0; i < 10000; i++ {
				wg.Add(1)
				go func() {
					cc.Inc()
					wg.Done()
				}()
			}

			wg.Wait()
		}
	})
}
