package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Cортировать можно не только инты
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// Постораемся реализовать улучшенный алгоритм быстрой сортировки
// qsort3. Он устраняет заметный недостаток обычной быстрой сортировки -
// крадратичное время выполнения на наборе данных, в котором очень много
// одинаковых элементов
func qsort3[T Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// refVal - опорный элемент
	// Будем выбирать опорный элемент случайным образом
	// Это поможет избежать сложностей с сортировкой
	// почти отсортированных данных
	refValue := arr[rand.Intn(len(arr))]

	// Less, Equal - индексы границ групп
	// После разделения на группы:
	// элементы с индексами 	от 0	 до Less-1		будут иметь значение меньше чем refVal
	// элементы с индексами 	от Less  до Equal-1 	будут иметь значение равное refVal
	// элементы с индексами 	от Equal до len(arr) 	будут иметь значение больше чем refVal
	//
	// Пример состояния слайса после разделения на группы с refVal = 10:
	//
	//                         Less       Equal
	//                           |          |
	// [2][0][5][4][6][4][0][5][10][10][10][15][11][25][12][30][33][34][55][99][11][17]
	Less, Equal := 0, 0

	// Разделение на группы
	for i := 0; i < len(arr); i++ {

		switch {
		// Случай первый: значение меньше опорного
		case arr[i] < refValue:

			// Разное поведение в зависимости от того, попадалось ли
			// значение равное опорному
			if Less == Equal {
				// текущее значение отправляется к группе меньших и встаёт последним
				arr[i], arr[Less] = arr[Less], arr[i]
			} else {
				// текущее значение отправляется к группе меньших и встаёт последним
				// чуть более сложным образом
				arr[i], arr[Equal] = arr[Equal], arr[i]
				arr[Less], arr[Equal] = arr[Equal], arr[Less]
			}

			// границы областей "< refVal" и "== refVal" сдвигаются
			Less++
			Equal++

		// Случай второй: значение равно опорному
		case arr[i] == refValue:
			// текущее значение отправляется к группе равных и встаёт последним
			arr[i], arr[Equal] = arr[Equal], arr[i]

			// граница области "== refVal" сдвигается
			Equal++
		}

		// В случае если текущее значение больше refVal ничего не делаем
		// просто идём дальше по массиву
	}

	// На данном этапе Элементы с индексами от Less ... до Equal-1
	// уже отсортированы. То есть занимают те же места что и в полностью отсортированном слайсе

	// Запустим сортировку рекурсивно на элементах, которые меньше опорного...
	qsort3(arr[:Less])

	// ... и на элементах больше опорного
	qsort3(arr[Equal:])

}

func main() {
	str := "日本語で何か言ってください。（にほんごでなにかいってください）A small language that compiles fast makes for a happy developer. The Go language is small, compiles really fast, and as a result it lets your mind focus on the actual problem and less on the tool you are using to solve it. Code, test, debug cycles are so quick that you forget you are not working with an interpreted language. Looking at our code, you see less boilerplate and more business logic. At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot! Реализовать быструю сортировку массива (quicksort) встроенными методами языка."
	runesToSort := []rune(str)

	fmt.Printf("Отсортируем руны в строке:\n\"%s\"\n", str)

	qsort3(runesToSort)

	myResult := string(runesToSort)

	runesToSort = []rune(str)

	fmt.Println("\nСравним с сортировкой из пакета sort...")

	sort.Slice(runesToSort, func(i, j int) bool { return runesToSort[i] < runesToSort[j] })

	goResult := string(runesToSort)

	fmt.Printf("(goResult == myResult): %v\n", (goResult == myResult))
}
