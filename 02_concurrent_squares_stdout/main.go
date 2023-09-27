package main

import (
	"fmt"
	"sync"
)

var (
	// в задании сказано "взятых из массива (2,4,6,8,10)"
	// если необходимо всё-таки использовать слайс, или
	// массив другой длинны, функция ConcurrentSquaresToStdout()
	// с этм справится. Необходимо лишь переопределить переменную array.
	array = [5]int{2, 4, 6, 8, 10}
)

func ConcurrentSquaresToStdout() {
	wg := &sync.WaitGroup{}

	for _, number := range array {
		wg.Add(1)

		// явно передавать number в функцию неоходмио из-за особенностей
		// переменных цикла for range {}
		go func(number int) {

			// функция fmt.Println() гарантировано пишет в os.Stdout
			// фрагмент её листинга:
			//
			// Fprintln(os.Stdout, a...)

			fmt.Println(number * number)
			wg.Done()
		}(number)
	}

	wg.Wait()
}

func main() {
	ConcurrentSquaresToStdout()
}
