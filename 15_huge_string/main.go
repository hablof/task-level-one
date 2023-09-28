package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unsafe"
)

var (
	justString string

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// В языке Go тип string является
// своего рода обёрткой над типом []byte
// который, в свою очередь, ссылается на массив байт фиксированной длины.
//
// Поэтому, когда функция createHugeString каким угодно образом создаёт большую строку,
// аллоцироется большой массив байт, а все производные строки лишь ссылаются на этот массив
// хотя и имеют разную длину.
//
// Продемонстрируем это, выведя unsafe.StringData(...) строк
//
// Такое переиспользование безопасно, ведь строки в Go иммутабельны.
//
// Таким образом негативные последствия использования представленного кода
// заключаются в том, что, и хотя строка justString имеет длину 100 байт,
// объём выделенной под неё памяти остаётся (1 << 10).
// То есть, нижележащий []byte имеел len(100) и cap(1024)

// Yиже представлены варианты обхода этого поведенния

func createHugeString(size int) string {
	byteSlice := make([]byte, size)
	for i := range byteSlice {
		byteSlice[i] = letters[rand.Intn(len(letters))]
	}

	return string(byteSlice)
}

func someFunc() {
	fmt.Println("Исходный вариант:")
	hugeString1 := createHugeString(1 << 10)
	fmt.Printf("pointer to hugeString1:\t%v, length of hugeString1:\t%d\n", unsafe.StringData(hugeString1), len(hugeString1))
	justString = hugeString1[:100]
	fmt.Printf("pointer to justString:\t%v, length of justString:\t%d\n\n", unsafe.StringData(justString), len(justString))
}

func main() {
	someFunc()

	sliceCopy()

	stringConcat()

	stringBuilder()

	stringsCopy()
}

// Один из вариантов решения проблемы
// строчкой newStringBytes := make([]byte, 100)
// мы явно выделяем новую память
func sliceCopy() {
	fmt.Println("Вариант с копированием в новый слайс:")
	hugeString2 := createHugeString(1 << 10)
	fmt.Printf("pointer to hugeString2:\t%v, length of hugeString2:\t%d\n", unsafe.StringData(hugeString2), len(hugeString2))

	newStringBytes := make([]byte, 100)

	copy(newStringBytes, hugeString2)

	justString = string(newStringBytes)
	fmt.Printf("pointer to justString:\t%v, length of justString:\t%d\n\n", unsafe.StringData(justString), len(justString))
}

// Ещё один способ спровоцировать выделение памяти
// Поскольку нет гарантии, что после конкатенации
// получится (часть) исходной строки,
// выделяется новая память
func stringConcat() {
	fmt.Println("Вариант с конкатенацией строк:")
	hugeString3 := createHugeString(1 << 10)
	fmt.Printf("pointer to hugeString3:\t%v, length of hugeString3:\t%d\n", unsafe.StringData(hugeString3), len(hugeString3))
	justString = hugeString3[:99] + hugeString3[99:100]
	fmt.Printf("pointer to justString:\t%v, length of justString:\t%d\n\n", unsafe.StringData(justString), len(justString))
}

// Этот вариант по сути повторяет вариант sliceCopy
// Но представлен чуть более дружелюбным синтаксисом
func stringBuilder() {
	fmt.Println("Вариант с strings.Builder:")
	hugeString4 := createHugeString(1 << 10)
	fmt.Printf("pointer to hugeString4:\t%v, length of hugeString4:\t%d\n", unsafe.StringData(hugeString4), len(hugeString4))

	sb := strings.Builder{}
	_, _ = sb.WriteString(hugeString4[:100])
	justString = sb.String()
	fmt.Printf("pointer to justString:\t%v, length of justString:\t%d\n\n", unsafe.StringData(justString), len(justString))
}

// Самый явный способ переаллоцировать строку за нас сделали авторы языка
// Это функция strings.Clone()
// По логике происходящего, внутри она работает точно так же как и наша функция sliceCopy()
// но поскольку используется unsafe.String(...) это позволяет избежать дополнительной аллокации
func stringsCopy() {
	fmt.Println("Вариант с слонированием строки:")
	hugeString5 := createHugeString(1 << 10)
	fmt.Printf("pointer to hugeString5:\t%v, length of hugeString5:\t%d\n", unsafe.StringData(hugeString5), len(hugeString5))

	justString = strings.Clone(hugeString5[:100])
	fmt.Printf("pointer to justString:\t%v, length of justString:\t%d\n\n", unsafe.StringData(justString), len(justString))
}
