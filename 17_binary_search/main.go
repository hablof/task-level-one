package main

import (
	"errors"
	"fmt"
)

// Cортировать можно не только инты
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

var (
	ErrNotFound = errors.New("target not found")
)

// Для бинарного поиска, слайс должен быть отсортирован
// Возвращает индекс элемента target или ошибку ErrNotFound,
// если такого элемента нет
func binarySearch[T Ordered](nums []T, target T) (int, error) {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2

	for left <= right {
		mid = left + (right-left)/2

		switch {
		case nums[mid] == target:
			return mid, nil

		case nums[mid] < target:
			left = mid + 1

		case nums[mid] > target:
			right = mid - 1
		}
	}

	return -1, ErrNotFound
}

func main() {
	arr := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}
	target := 90

	fmt.Printf("В отсортированном слайсе\n%v\nНайдём индекс элемента %d\n...", arr, target)

	idx, err := binarySearch(arr, target)
	if err != nil {
		fmt.Printf("Не удалось найти элемент %d\n", target)
	} else {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, idx)
	}

	target = 15
	fmt.Printf("Попытаемся найти индекс элемента %d...\n", target)

	idx, err = binarySearch(arr, target)
	if err != nil {
		fmt.Printf("Не удалось найти элемент %d\n", target)
	} else {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, idx)
	}
}
