package main

import (
	"fmt"
	"sync"
)

// знаю три способа обеспечить безопасность конкурентной записи в map:

// 1) использование Mutex
type myMapMutex struct {
	mu sync.Mutex
	m  map[int]int
}

func (m *myMapMutex) ConcurrenSafeWrite(k, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}

// 2) использование RWMutex
type myMapRWMutex struct {
	mu sync.RWMutex
	m  map[int]int
}

func (m *myMapRWMutex) ConcurrenSafeWrite(k, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}

// 3) использование SyncMap
type myMapSyncMap struct {
	m *sync.Map
}

func (m *myMapSyncMap) ConcurrenSafeWrite(k, v int) {
	m.m.Store(k, v)
}

func main() {
	mapSyncMap := myMapSyncMap{
		m: &sync.Map{},
	}

	mapMutex := myMapMutex{
		mu: sync.Mutex{},
		m:  map[int]int{},
	}

	mapRWMutex := myMapRWMutex{
		mu: sync.RWMutex{},
		m:  map[int]int{},
	}

	wg := &sync.WaitGroup{}

	// десять раз по одинаковым ключам
	for i := 0; i < 10; i++ {
		for j := 0; j < 1000; j++ {
			wg.Add(3)

			go func(i int) {
				defer wg.Done()
				mapSyncMap.ConcurrenSafeWrite(i, i)
			}(j)

			go func(i int) {
				defer wg.Done()
				mapMutex.ConcurrenSafeWrite(i, i)
			}(j)

			go func(i int) {
				defer wg.Done()
				mapRWMutex.ConcurrenSafeWrite(i, i)
			}(j)
		}
	}

	wg.Wait()
	fmt.Printf("len(mapMutex.m): %v\n", len(mapMutex.m))
	fmt.Printf("len(mapRWMutex.m): %v\n", len(mapRWMutex.m))

	counter := 0
	mapSyncMap.m.Range(func(key, value any) bool {
		counter++
		return true
	})

	fmt.Printf("len(mapSyncMap.m): %v\n", len(mapRWMutex.m))
}
