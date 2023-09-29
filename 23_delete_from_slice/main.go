package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotInBounds = errors.New("index out of length")
)

// Функция удаляет элемент под индексом idx и сдвигает все элементы следующие за ним
// Сложность линейная O(len(slice) - idx)
//
// Переаллокация не происходит
// Занимаемый объём памяти нижележащим массивом не уменьшается
func DeleteIthFromSliceWithSavedOrder[T any](slice *[]T, idx int) error {
	if idx >= len(*slice) || idx < 0 {
		return ErrNotInBounds
	}

	*slice = append((*slice)[:idx], (*slice)[idx+1:]...)

	return nil
}

// Функция удаляет элемент под индексом idx следующим образом:
// 1) элемент под индексом idx меняется местами с последним элементом
// 2) последний элемент отрезается от среза
// Сложность константная O(1)
//
// Переаллокация не происходит
// Занимаемый объём памяти нижележащим массивом не уменьшается
func DeleteIthFromSliceNoOrder[T any](slice *[]T, idx int) error {
	if idx >= len(*slice) || idx < 0 {
		return ErrNotInBounds
	}

	(*slice)[len(*slice)-1], (*slice)[idx] = (*slice)[idx], (*slice)[len(*slice)-1]

	*slice = (*slice)[:len(*slice)-1]

	return nil
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	refPointer := &arr[0]

	fmt.Printf("arr: %v, len(arr)==%d, cap(arr)==%d\n", arr, len(arr), cap(arr))
	fmt.Printf("Из массива: %v\nудалим первый элемент с сохранением порядка...\n", arr)
	DeleteIthFromSliceWithSavedOrder(&arr, 0)

	fmt.Printf("Из массива: %v\nудалим первый элемент без сохранения порядка...\n", arr)
	DeleteIthFromSliceNoOrder(&arr, 0)

	fmt.Printf("Из массива: %v\nудалим десятый элемент с сохранением порядка...\n", arr)
	DeleteIthFromSliceWithSavedOrder(&arr, 0)

	fmt.Printf("Из массива: %v\nудалим десятый элемент без сохранения порядка...\n", arr)
	DeleteIthFromSliceNoOrder(&arr, 0)

	fmt.Printf("arr: %v, len(arr)==%d, cap(arr)==%d\n", arr, len(arr), cap(arr))

	fmt.Printf("(refPointer == &arr[0]): %v\n", (refPointer == &arr[0]))
}
