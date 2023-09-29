package concat

import (
	"fmt"
	"strings"
)

// В этом вопросе нужно рассмотреть два принципиально разных случая

// 1) количество строк фиксированно и известно заранее, ещё на этапе компиляции
//    напрмер, структура имеет 5 строковых полей, и нужно их проконкатенировать

// 2) количество строк НЕ известно заранее и НЕ фиксированно
//    например - массив строк пришедший по сети в формате json

// Результаты бенчмарков на моей машине:

// goos: windows
// goarch: amd64
// pkg: github.com/hablof/task-level-one/questions/01_concat
// cpu: Intel(R) Core(TM) i3-2100 CPU @ 3.10GHz

// Benchmark_Fixed/concatStrings-4         	 7335823	       175.0 ns/op	     208 B/op	       1 allocs/op
// Benchmark_Fixed/formatString-4          	 1678976	       716.2 ns/op	     304 B/op	       7 allocs/op
// Benchmark_Fixed/Builder_NoPreAlloc-4    	 3227091	       378.6 ns/op	     472 B/op	       4 allocs/op
// Benchmark_Fixed/Builder_PreAlloc-4      	 7297849	       163.1 ns/op	     208 B/op	       1 allocs/op

// Benchmark_NonFixed/concatStrings-4      	   56222	     20048 ns/op	   34376 B/op	      95 allocs/op
// Benchmark_NonFixed/formatString-4       	  132216	      9889 ns/op	    4034 B/op	      99 allocs/op
// Benchmark_NonFixed/Builder_NoPreAlloc-4 	 1000000	      2132 ns/op	    2600 B/op	       8 allocs/op
// Benchmark_NonFixed/Builder_PreAlloc-4   	 1000000	      1140 ns/op	     768 B/op	       1 allocs/op

// Вывод: самым эффективным методом конкатенации в обоих случаях является использование
// strings.Builder с предварительным выделением памяти.
// Однако заметим, что в случае фиксированного количества строк, прямая конкатенация
// совсем незначительно уступает как в быстродействии, так и в использовании памяти

// 1) реализуем различные варианты конкатенации фиксировнного количества строк

// прямая конкатенация
func concatStrings_Fixed(s1, s2, s3, s4, s5, s6 string) string {
	return s1 + s2 + s3 + s4 + s5 + s6
}

// подставление в форматирущую строку
func formatString_Fixed(s1, s2, s3, s4, s5, s6 string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", s1, s2, s3, s4, s5, s6)
}

// использование strings.Builder без предварительного выделения памяти
func stringsBuilder_Fixed_NoPreAlloc(s1, s2, s3, s4, s5, s6 string) string {
	var sb strings.Builder

	sb.WriteString(s1)
	sb.WriteString(s2)
	sb.WriteString(s3)
	sb.WriteString(s4)
	sb.WriteString(s5)
	sb.WriteString(s6)

	return sb.String()
}

// использование strings.Builder с предварительным выделением памяти
func stringsBuilder_Fixed_PreAlloc(s1, s2, s3, s4, s5, s6 string) string {
	var sb strings.Builder

	sb.Grow(len(s1) + len(s2) + len(s3) + len(s4) + len(s5) + len(s6))
	sb.WriteString(s1)
	sb.WriteString(s2)
	sb.WriteString(s3)
	sb.WriteString(s4)
	sb.WriteString(s5)
	sb.WriteString(s6)

	return sb.String()
}

//
// 2) реализуем различные варианты конкатенации нефиксированого количества строк
//

// прямая конкатенация в цикле
func concatStrings_NonFixed(s1 ...string) string {
	var s string
	for _, ss := range s1 {
		s = s + ss
	}

	return s
}

// создание форматированной строки и передача туда аргументов
func formatString_NonFixed(s1 ...string) string {
	sliceAny := make([]any, 0, len(s1))

	for _, v := range s1 {
		sliceAny = append(sliceAny, v)
	}

	template := strings.Repeat("%s", len(s1))

	return fmt.Sprintf(template, sliceAny...)
}

// использование strings.Builder без предварительного выделения памяти
func stringsBuilder_NonFixed_NoPreAlloc(s1 ...string) string {
	var sb strings.Builder

	for _, ss := range s1 {
		sb.WriteString(ss)
	}

	return sb.String()
}

// использование strings.Builder с предварительным выделением памяти
func stringsBuilder_NonFixed_PreAlloc(s1 ...string) string {
	var sb strings.Builder

	otputSize := 0
	for _, ss := range s1 {
		otputSize += len(ss)
	}
	sb.Grow(otputSize)

	for _, ss := range s1 {
		sb.WriteString(ss)
	}

	return sb.String()
}
