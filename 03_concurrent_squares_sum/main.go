package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Не совсем ясно в каком виде будет дана "последовательность"
// но в функцию с множественным агрументом можно передать
// и неограниченное количество отдельных аргументов
// и слайс (синтаксис: func(slice...))
// и массив предварительно взяв от него слайс (синтаксис: func(arr[:]...))

// Вариант функции с мьютексом,
// который защищает вычисление квадрата и операцию "+="
func SumSquaresMutex(numbers ...int) int {

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	accumulator := 0

	for _, number := range numbers {
		wg.Add(1)

		go func(number int) {
			mu.Lock()
			accumulator += number * number
			mu.Unlock()

			wg.Done()
		}(number)
	}

	wg.Wait()

	return accumulator
}

// Вариант функции с каналом
// В этом варианте нет переменных,
// к которым имеется доступ из нескольких горутин
// соответственно мьютекс не нужен, состояние гонки невозможно
// также, возможно отказаться от WaitGroup, читая из канала фиксированное количество раз
func SumSquaresChannel(numbers ...int) int {
	calculationChannel := make(chan int)

	for _, number := range numbers {
		go func(number int) {
			calculationChannel <- number * number
		}(number)
	}

	accumulator := 0
	for i := 0; i < len(numbers); i++ {
		accumulator += <-calculationChannel
	}

	return accumulator
}

// то же что и SumSquaresChannel, но с более элегантным выходом из чтения
func SumSquaresChannelForRange(numbers ...int) int {
	calculationChannel := make(chan int)
	wg := sync.WaitGroup{}

	for _, number := range numbers {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			calculationChannel <- number * number
		}(number)
	}

	go func() {
		wg.Wait()
		close(calculationChannel)
	}()

	accumulator := 0
	for num := range calculationChannel {
		accumulator += num
	}

	return accumulator
}

// Вариант функции с атомарными переменными,
func SumSquaresAtomic(numbers ...int) int {

	accumulator := atomic.Int64{}
	wg := &sync.WaitGroup{}

	for _, number := range numbers {
		wg.Add(1)

		go func(number int) {
			accumulator.Add(int64(number * number))
			wg.Done()
		}(number)
	}

	wg.Wait()

	return int(accumulator.Load())
}

var (
	numbers = []int{2, 4, 6, 8, 10}
)

func main() {
	fmt.Printf("SumSquaresAtomic(): %v\n", SumSquaresAtomic(numbers...))
	fmt.Printf("SumSquaresChannel(): %v\n", SumSquaresChannel(numbers...))
	fmt.Printf("SumSquaresChannelForRange(): %v\n", SumSquaresChannelForRange(numbers...))
	fmt.Printf("SumSquaresMutex(): %v\n", SumSquaresMutex(numbers...))
}
