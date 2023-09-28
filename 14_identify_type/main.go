package main

import (
	"fmt"
	"reflect"
)

var (
	arr = []interface{}{15, "str", true, make(chan struct{}), make(chan chan struct{})}
)

// можно определять тип с помощью конструкции type switch
// минусом такого подхода является необходимость заранее знать,
// какой тип может иметь переданное значение
// особенно ярко \то видно на строчке "case chan struct{}:".
// Однако конструкция type switch позволяет ветвить логику программы
// и использовать приведённое к типу значение.
func TypeSwitch(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"

	case string:
		return "string"

	case bool:
		return "bool"

	case chan struct{}:
		return "chan struct {}"
	}

	return "unknown type"
}

// более простой способ решения этой задачи - использовать плейсхолдер %T
// однако, помимо строки с типом этот подход не позволяет получить ничего
func TypeSprintf(val interface{}) string {
	return fmt.Sprintf("%T", val)
}

// плейсхолдер %T внутри себя использует пакет reflect
// мы тоже можем его использовать
func TypeReflect(val interface{}) string {
	return reflect.TypeOf(val).String()
}
func main() {
	for _, elem := range arr {
		fmt.Printf("type of %v: %s\t", elem, TypeSwitch(elem))
	}
	fmt.Println()

	for _, elem := range arr {
		fmt.Printf("type of %v: %s\t", elem, TypeSprintf(elem))
	}
	fmt.Println()

	for _, elem := range arr {
		fmt.Printf("type of %v: %s\t", elem, TypeReflect(elem))
	}
}
