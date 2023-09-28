package main

import (
	"fmt"
)

func ReadNumbersFromArray(arr []int) <-chan int {
	outputCh := make(chan int)

	// Последовательно читает числа из массива, пишет их в канал
	// После того как все числа записаны в канал, закрывает его и завершается
	go func() {
		for _, elem := range arr {
			outputCh <- elem
		}

		close(outputCh)
	}()

	return outputCh
}

func SquareNumbers(inputCh <-chan int) <-chan int {
	outputCh := make(chan int)

	// Читает из какнала приходящие числа и тут же записывает
	// в другой канал их квадраты
	// После того как канал inputCh будет закрыт другой горутиной,
	// перестаёт читать и закрывает канал outputCh
	go func() {
		for elem := range inputCh {
			outputCh <- elem * elem
		}

		close(outputCh)
	}()

	return outputCh
}

func DoubleNumbers(inputCh <-chan int) <-chan int {
	outputCh := make(chan int)

	// Читает из канала приходящие числа и тут же записывает
	// в другой канал их удвоенное значение
	// После того как канал inputCh будет закрыт другой горутиной,
	// перестаёт читать и закрывает канал outputCh
	go func() {
		for elem := range inputCh {
			outputCh <- elem * 2
		}

		close(outputCh)
	}()

	return outputCh
}

// Читает числа из канала и выводит в stdout
// После закрытия канала завершается
func PrintNumbers(inputCh <-chan int) {
	for elem := range inputCh {
		fmt.Printf("%d ", elem)
	}
}

var (
	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {

	chanWithNumbers := ReadNumbersFromArray(arr)

	// Запускаем в горутине ф-ю умножающую числа
	//
	// в задании написано "результат операции x*2"
	// но я заподозрил очепятку: возможно, имелось в виду "x^2"
	// тогда надо использовать функцию SquareNumbers

	// chanWithSquares := SquareNumbers(chanWithNumbers)
	chanWithSquares := DoubleNumbers(chanWithNumbers)

	PrintNumbers(chanWithSquares)
}
