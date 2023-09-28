package main

import (
	"fmt"

	"github.com/hablof/task-level-one/11_sets_intersect/sets"
)

var (
	// последовательность содержит этинаковые элементы
	// а множество не может сдержать повторяющиеся элементы
	// см. https://ru.wikipedia.org/wiki/Множество#Элемент_множества
	// поэтому в полученном множестве будет всего 3 элемента
	// "dog", "cat" и "tree"
	sequence = []string{"cat", "cat", "dog", "cat", "tree"}
)

func main() {
	// функциональность множеств реализована в предыдущем задании
	// создаём множество из слайса squence
	setOfStrings := sets.NewSet(sequence...)

	// выводим в stdout его элементы
	fmt.Printf("Элементы множества: %v\n", setOfStrings.Elements())
}
