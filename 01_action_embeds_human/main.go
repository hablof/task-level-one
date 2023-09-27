package main

import "fmt"

type Human struct {
	humanField1 int
	humanField2 int
	humanField3 string
}

func (h *Human) ChangeHumanFields() {
	h.humanField1, h.humanField2 = h.humanField1+h.humanField2, h.humanField1-h.humanField2
}

func (h *Human) HumanInfo() string {
	return fmt.Sprintf("human info: %d:%s:%d", h.humanField1, h.humanField3, h.humanField2)
}

type Action struct {
	actionField1 string
	actionField2 string

	// Данный синтаксис позволяет реализовать "встраивание" в golang.
	// Теперь тип имеет и встроенные поля humanField1..3 и имеет методы
	// .ChangeHumanFields() и .HumanInfo()
	Human
}

func (a *Action) TakeAction() {
	fmt.Printf("\"nice fields: %s, %s, %s\"\n", a.humanField3, a.actionField1, a.actionField2)
}

func main() {
	// Создадим экземпляр структуры Action
	a := Action{
		actionField1: "a1",
		actionField2: "a2",
		Human: Human{
			humanField1: 1,
			humanField2: 2,
			humanField3: "h3",
		},
	}

	// продемонстрируем работу методов самой структуры Action
	// и методы встроенной структуры Human структуры Action
	a.TakeAction()
	fmt.Printf("a.HumanInfo(): %v\n", a.HumanInfo())

	fmt.Println("a.ChangeHumanFields()...")
	a.ChangeHumanFields()

	fmt.Printf("a.HumanInfo(): %v\n", a.HumanInfo())

	// в частности, обратимся к полям встроенной структуры напрямую
	fmt.Printf("a.humanField1: %v\n", a.humanField1)
	fmt.Println("Increment a.humanField1...")
	a.humanField1++
	fmt.Printf("a.humanField1: %v\n", a.humanField1)
}
