package adapter

type GoConsole interface {
	Print(a ...any)
	Scan(a ...any)
}

// Адаптер применяется в качестве промежуточного звена между необходимым интерфесом
// и реализацией этого интерфеса с неподходящими названиями функций
type Adapter struct {
	// можно было обойтись и без интерфейса GoConsole
	// а просто напрямую встроить структуру goconsole.Console{}
	// в структуру Adapter
	//
	// goconsole.Console

	// Поле экспортируемое чтобы не писать функцию NewAdapter(gc GoConsole) *Adapter {}
	GoConsole GoConsole
}

// Вызовы методов Адаптера просто мапятся в вызовы методов включённого интерфейса.
func (ad *Adapter) WriteLn(a ...any) {
	ad.GoConsole.Print(a...)
}

func (ad *Adapter) ReadLn(a ...any) {
	ad.GoConsole.Scan(a...)
}
