package main

import "fmt"

func main() {
	h := Human{
		age:  33,
		name: "Testik",
	}
	fmt.Println("Создали структуру", h)

	action := Action{
		h,
		"какое-то особое действие",
	}

	// Можем вызвать DoAction() и Say() т.к. мы организовали встраивание
	action.DoAction() // Метод из структуры Human
	action.Say()      // Метод из структуры Action
}

type Human struct {
	name string
	age  int
}

func (h *Human) Say() {
	fmt.Printf("Hello, my name is %s\n", h.name)
}

type Action struct {
	Human
	ActionDescription string
}

func (a *Action) DoAction() {
	fmt.Printf("Person %s action: %s\n", a.name, a.ActionDescription)
}
