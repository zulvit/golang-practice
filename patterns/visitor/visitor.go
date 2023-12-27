package visitor

import "fmt"

// ComputerPart представляет интерфейс компонента компьютера
type ComputerPart interface {
	Accept(visitor ComputerPartVisitor)
}

// ComputerPartVisitor представляет интерфейс посетителя
type ComputerPartVisitor interface {
	VisitForMouse(mouse *Mouse)
	VisitForKeyboard(keyboard *Keyboard)
	VisitForMonitor(monitor *Monitor)
	VisitForComputer(computer *Computer)
}

// Mouse представляет компонент мыши
type Mouse struct{}

func (m *Mouse) Accept(visitor ComputerPartVisitor) {
	visitor.VisitForMouse(m)
}

// Keyboard представляет компонент клавиатуры
type Keyboard struct{}

func (k *Keyboard) Accept(visitor ComputerPartVisitor) {
	visitor.VisitForKeyboard(k)
}

// Monitor представляет компонент монитора
type Monitor struct{}

func (m *Monitor) Accept(visitor ComputerPartVisitor) {
	visitor.VisitForMonitor(m)
}

// Computer представляет компьютер
type Computer struct {
	Parts []ComputerPart
}

func (c *Computer) Accept(visitor ComputerPartVisitor) {
	for _, part := range c.Parts {
		part.Accept(visitor)
	}
	visitor.VisitForComputer(c)
}

func NewComputer() *Computer {
	return &Computer{
		Parts: []ComputerPart{&Mouse{}, &Keyboard{}, &Monitor{}},
	}
}

// ComputerPartDisplayVisitor представляет посетителя, который отображает компоненты
type ComputerPartDisplayVisitor struct{}

func (v *ComputerPartDisplayVisitor) VisitForMouse(mouse *Mouse) {
	fmt.Println("Displaying Mouse.")
}

func (v *ComputerPartDisplayVisitor) VisitForKeyboard(keyboard *Keyboard) {
	fmt.Println("Displaying Keyboard.")
}

func (v *ComputerPartDisplayVisitor) VisitForMonitor(monitor *Monitor) {
	fmt.Println("Displaying Monitor.")
}

func (v *ComputerPartDisplayVisitor) VisitForComputer(computer *Computer) {
	fmt.Println("Displaying Computer.")
}
