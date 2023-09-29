package mutex

import (
	"fmt"
	"sync"
	"testing"
)

// получил довольно неожиданный для себя результат
// goos: windows
// goarch: amd64
// pkg: github.com/hablof/task-level-one/questions/03_RWMutex_Mutex
// cpu: Intel(R) Core(TM) i3-2100 CPU @ 3.10GHz
// Benchmark_Mutexes/R:W_ratio___1/stMutex-4 	     133	   8885380 ns/op	  488043 B/op	   20025 allocs/op
// Benchmark_Mutexes/R:W_ratio___1/RWMutex-4 	     121	   9694017 ns/op	  483168 B/op	   20035 allocs/op
// Benchmark_Mutexes/R:W_ratio___2/stMutex-4 	     122	   9820641 ns/op	  480555 B/op	   20009 allocs/op
// Benchmark_Mutexes/R:W_ratio___2/RWMutex-4 	     123	   9637846 ns/op	  481581 B/op	   20018 allocs/op
// Benchmark_Mutexes/R:W_ratio___4/stMutex-4 	     121	   9651040 ns/op	  480457 B/op	   20007 allocs/op
// Benchmark_Mutexes/R:W_ratio___4/RWMutex-4 	     123	   9821673 ns/op	  484759 B/op	   20047 allocs/op
// Benchmark_Mutexes/R:W_ratio__10/stMutex-4 	     100	  10075909 ns/op	  480935 B/op	   20012 allocs/op
// Benchmark_Mutexes/R:W_ratio__10/RWMutex-4 	     100	  10294721 ns/op	  491129 B/op	   20114 allocs/op
// Benchmark_Mutexes/R:W_ratio__25/stMutex-4 	     120	  11212274 ns/op	  482634 B/op	   20030 allocs/op
// Benchmark_Mutexes/R:W_ratio__25/RWMutex-4 	     121	   9877838 ns/op	  487309 B/op	   20078 allocs/op
// Benchmark_Mutexes/R:W_ratio_100/stMutex-4 	     123	   9829226 ns/op	  480598 B/op	   20008 allocs/op
// Benchmark_Mutexes/R:W_ratio_100/RWMutex-4 	     124	   9617587 ns/op	  481466 B/op	   20017 allocs/op
// PASS
// ok  	github.com/hablof/task-level-one/questions/03_RWMutex_Mutex	23.992s
//
// RWMutex начинает себя лучше показывать только начиная с соотношения чтение/запись больше 25
// При этом отноительная разница составляет единицы процентов
// От запуска к запуску результаты разнятся :(
type simpleMutex struct {
	mu      *sync.Mutex
	counter int
}

func (sm *simpleMutex) Read() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	return sm.counter
}

func (sm *simpleMutex) Write() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.counter++
}

type rwMutex struct {
	mu      *sync.RWMutex
	counter int
}

func (rw *rwMutex) Read() int {
	rw.mu.RLock()
	defer rw.mu.RUnlock()

	return rw.counter
}

func (rw *rwMutex) Write() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.counter++
}

func Benchmark_Mutexes(b *testing.B) {
	concurrentActions := 20000
	ratiosNumerator := []int{1, 2, 4, 10, 25, 100}

	for ratio := range ratiosNumerator {
		b.Run(fmt.Sprintf("R:W_ratio_%3d", ratiosNumerator[ratio]), func(b *testing.B) {

			b.Run("S_Mutex", func(b *testing.B) { // S for standart
				for i := 0; i < b.N; i++ {

					s := simpleMutex{
						mu:      &sync.Mutex{},
						counter: 0,
					}

					wg := &sync.WaitGroup{}
					actions := 0
					run := 0

					for actions < concurrentActions {
						run++
						wg.Add(1)
						actions++

						go func() {
							_ = s.Read()
							wg.Done()
						}()

						if run%ratiosNumerator[ratio] == 0 {
							wg.Add(1)
							actions++

							go func() {
								s.Write()
								wg.Done()
							}()
						}
					}

					wg.Wait()
				}
			})

			b.Run("RWMutex", func(b *testing.B) {
				for i := 0; i < b.N; i++ {

					s := rwMutex{
						mu:      &sync.RWMutex{},
						counter: 0,
					}

					wg := &sync.WaitGroup{}
					actions := 0
					run := 0

					for actions < concurrentActions {
						wg.Add(1)
						actions++

						go func() {
							_ = s.Read()
							wg.Done()
						}()

						if run%ratiosNumerator[ratio] == 0 {
							wg.Add(1)
							actions++

							go func() {
								s.Write()
								wg.Done()
							}()
						}
					}

					wg.Wait()
				}
			})
		})
	}

}
