package goconsole

import "fmt"

// Тип, который мы будем адаптировать
type Console struct{}

func (Console) Print(a ...any) {
	fmt.Println(a...)
}

func (Console) Scan(a ...any) {
	fmt.Scanln(a...)
}
