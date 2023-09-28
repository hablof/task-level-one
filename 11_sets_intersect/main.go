package main

import (
	"fmt"

	"github.com/hablof/task-level-one/11_sets_intersect/sets"
)

var (
	floats1 = []float64{0.1, 155.444, -4000.0004}
	floats2 = []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}

	ints1 = []int{1, 2, 3, 4, 5}
	ints2 = []int{6, 7, 8, 9, 10}

	strings1 = []string{"rofl", "kek", "lol", "heh", "ахахахах"}
	strings2 = []string{"rofl", "kek", "lol", "heh", "ахах"}
)

func main() {
	s1 := sets.NewSet(floats1...)
	s2 := sets.NewSet(floats2...)
	fmt.Println("Пересечём два множества чисел с плавающей точкой")
	fmt.Printf("Элементы первого множества:    %v\n", s1.Elements())
	fmt.Printf("Элементы второго множества:    %v\n", s2.Elements())
	fmt.Printf("Элементы пересечения множеств: %v\n", s2.Intersect(s1).Elements())

	fmt.Println()

	s3 := sets.NewSet(ints1...)
	s4 := sets.NewSet(ints2...)
	fmt.Println("Пересечём два множества целых чисел")
	fmt.Printf("Элементы первого множества:    %v\n", s3.Elements())
	fmt.Printf("Элементы второго множества:    %v\n", s4.Elements())
	fmt.Printf("Элементы пересечения множеств: %v\n", s3.Intersect(s4).Elements())

	fmt.Println()

	s5 := sets.NewSet(strings1...)
	s6 := sets.NewSet(strings2...)
	fmt.Println("Пересечём два множества строк")
	fmt.Printf("Элементы первого множества:    %v\n", s5.Elements())
	fmt.Printf("Элементы второго множества:    %v\n", s6.Elements())
	fmt.Printf("Элементы пересечения множеств: %v\n", s6.Intersect(s5).Elements())

}
