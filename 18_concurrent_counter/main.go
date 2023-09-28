package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализация конкуррентного счётчика на атомарных переменных
type atomicCounter struct {
	counter *atomic.Uint64
}

func (ac *atomicCounter) Inc() {
	ac.counter.Add(1)
}

func (ac *atomicCounter) Value() uint64 {
	return ac.counter.Load()
}

// Реализация конкуррентного счётчика через мьютекс
// В моём локальном запуске бенчмарков разница в производительности незначительна:
// goos: windows
// goarch: amd64
// pkg: github.com/hablof/task-level-one/task_18
// cpu: Intel(R) Core(TM) i3-2100 CPU @ 3.10GHz
// Benchmark_Counter/mutex_counter-4	253	   4455636 ns/op	245481 B/op		10019 allocs/op
// Benchmark_Counter/atomic_counter-4	288	   4611360 ns/op	240016 B/op		10001 allocs/op
type mutexCounter struct {
	mu      *sync.Mutex
	counter uint64
}

func (mc *mutexCounter) Inc() {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	mc.counter++
}

func (mc *mutexCounter) Value() uint64 {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	return mc.counter
}

// Должен быть создан функцией newChannelCounter()
// попытка испльзовать неинициализированную структуру
// приведёт к безвозвратной блокировке горутины
type channelCounter struct {
	incChn  chan struct{}
	outChn  chan uint64
	counter uint64
}

func newChannelCounter(ctx context.Context) *channelCounter {
	cc := &channelCounter{
		incChn:  make(chan struct{}),
		outChn:  make(chan uint64),
		counter: 0,
	}

	go func() {
		for {
			select {
			case <-cc.incChn:
				cc.counter++
			case cc.outChn <- cc.counter:

			case <-ctx.Done():
				return
			}
		}
	}()

	return cc
}

func (cc *channelCounter) Inc() {
	cc.incChn <- struct{}{}
}

func (cc *channelCounter) Value() uint64 {
	return <-cc.outChn
}

func main() {
	mCounter := mutexCounter{
		mu:      &sync.Mutex{},
		counter: 0,
	}

	aCounter := atomicCounter{
		counter: &atomic.Uint64{},
	}

	ctx, cf := context.WithCancel(context.Background())
	defer cf()

	cCounter := newChannelCounter(ctx)

	defer func() {
		fmt.Printf("cc.Value(): %v\n", cCounter.Value())
	}()
	defer func() {
		fmt.Printf("aCounter.Value(): %v\n", aCounter.Value())
	}()
	defer func() {
		fmt.Printf("mCounter.Value(): %v\n", mCounter.Value())
	}()

	wg := &sync.WaitGroup{}

	wg.Add(30000)
	fmt.Printf("Проинкрементируем счётчики конкуррентным образом\n")
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			aCounter.Inc()
		}()
		go func() {
			defer wg.Done()
			mCounter.Inc()
		}()
		go func() {
			defer wg.Done()
			cCounter.Inc()
		}()
	}

	wg.Wait()
}
