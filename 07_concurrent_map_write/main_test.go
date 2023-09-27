package main

import (
	"sync"
	"testing"
)

func Test_ConcurrentSafeWrite(t *testing.T) {
	t.Run("test mutex", func(t *testing.T) {
		m := myMapMutex{
			mu: sync.Mutex{},
			m:  map[int]int{},
		}

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(i, i)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(1000-i, i)
		}

		wg.Wait()
		t.Log("mutex finished")
	})

	t.Run("test RWmutex", func(t *testing.T) {
		m := myMapRWMutex{
			mu: sync.RWMutex{},
			m:  map[int]int{},
		}

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(i, i)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(1000-i, i)
		}

		wg.Wait()
		t.Log("RWmutex finished")
	})

	t.Run("test syncMap", func(t *testing.T) {
		m := myMapSyncMap{
			m: &sync.Map{},
		}

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(i, i)
			go func(k, v int) {
				defer wg.Done()
				m.ConcurrenSafeWrite(k, v)
			}(1000-i, i)
		}

		wg.Wait()
		t.Log("syncMap finished")
	})
}
