package main

import (
	"fmt"
	"sort"
)

var (
	temperatureOscillations = []float32{-1.7, 5.25, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
)

// Не согласен с представленном в задании примере распределения
// значений по группам
//
// В моей программе значения распределяются по группам ЕДИНЫМ образом
// Значение относится к группе десятка, который МЕНЬШЕ значения
// т.е. значения -25.0, -27.0, -21.0 будут объединены в группу "-30"
// иначе не ясно, какой группе должно принадлежать число -1.0
func determineGroup(val float32) int {
	group := (int(val) / 10) * 10
	if val < 0 {
		group -= 10
	}
	return group
}

func main() {
	fmt.Printf("Последовательность температурных колебаний: %v\n", temperatureOscillations)
	fmt.Println("Разобьём значения по группам с шагом в 10:")
	// мапа отображающая ключ группы на слайс значений этой группы
	temperatureGroups := make(map[int][]float32)

	// проходимся по значениям температур
	for _, val := range temperatureOscillations {
		// определяем какой группе значение должно принадлежать
		groupIdx := determineGroup(val)
		// добавляем значение в слайс группы
		temperatureGroups[groupIdx] = append(temperatureGroups[groupIdx], val)
	}

	// сказано, что последовательность в подмножествах не важна
	// но, возможно, важна последовательность самих групп
	groupNames := make([]int, 0, len(temperatureGroups))
	for groupName := range temperatureGroups {
		groupNames = append(groupNames, groupName)
	}

	// дя этого отсортируем ключи групп
	sort.Ints(groupNames)

	// выведем группы в правильном порядке
	for _, groupName := range groupNames {
		fmt.Printf("%3d: %v\n", groupName, temperatureGroups[groupName])
	}
}
