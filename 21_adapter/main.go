package main

import (
	"strconv"

	"github.com/hablof/task-level-one/21_adapter/adapter"
	goconsole "github.com/hablof/task-level-one/21_adapter/go_console"
)

// Интерфейс, под которым мы будем адаптировать
type PascalConsole interface {
	WriteLn(a ...any)
	ReadLn(a ...any)
}

func newPascalConsole(pc PascalConsole) PascalConsole {
	return pc
}

func main() {
	gc := goconsole.Console{}

	ad := adapter.Adapter{
		GoConsole: gc,
	}
	pc := newPascalConsole(&ad)

	r, g, b := 0, 0, 0

	pc.WriteLn("Введите значения для R, G и B")
	pc.ReadLn(&r, &g, &b)

	if r > 255 || g > 255 || b > 255 {
		pc.WriteLn("Значения не могут быть больше 255")
	}

	pc.WriteLn("ваш hex:", rgbToHex(r, g, b))
}

func rgbToHex(r, g, b int) string {
	return "#" + strconv.FormatInt(int64(1<<24|r<<16|g<<8|b), 16)[1:]
}
