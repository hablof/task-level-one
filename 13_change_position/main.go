package main

import "fmt"

func ChangePositions(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b

	// Пример: *a = 10, *b = 5
	// 1) *a = 10 + 5  | *a = 15, *b = 5
	// 2) *b = 15 - 5  | *a = 15, *b = 10
	// 3) *a = 15 - 10 | *a = 5, *b = 10
}

func ChangePositionsGolangStyle(a, b *int) {
	// это работает следующим оразом:
	// сначала последовательно вычисляются значения в правой части выражения
	// а зетем происходит присваивание значений переменным
	*a, *b = *b, *a

	// Пример: *a = 10, *b = 5
	// 1) *a, *b = *b, *a
	// 2) *a, *b = 5, 10
	// 3) *a = 5, *b = 10
}

func main() {
	a, b := 25, 14

	fmt.Printf("a: %d, b: %d\n", a, b)

	fmt.Println("changing positions...")
	ChangePositions(&a, &b)

	fmt.Printf("a: %d, b: %d\n", a, b)

	fmt.Println("changing positions...")
	ChangePositionsGolangStyle(&a, &b)

	fmt.Printf("a: %d, b: %d\n", a, b)

}
