package main

import "fmt"

type State interface {
	Write(text string) string
}

type SelectionState struct{}

func (s *SelectionState) Write(text string) string {
	return fmt.Sprintf("Текст \"%s\" выделен", text)
}

type TypingState struct{}

func (s *TypingState) Write(text string) string {
	return fmt.Sprintf("Введен текст: \"%s\"", text)
}

type ViewingState struct{}

func (s *ViewingState) Write(text string) string {
	return "Просмотр текста"
}

type TextEditor struct {
	state State
}

func (e *TextEditor) SetState(state State) {
	e.state = state
}

func (e *TextEditor) Type(text string) string {
	return e.state.Write(text)
}

func main() {
	editor := TextEditor{}

	editor.SetState(&TypingState{})
	fmt.Println(editor.Type("Привет, мир!"))

	editor.SetState(&SelectionState{})
	fmt.Println(editor.Type("Привет, мир!"))

	editor.SetState(&ViewingState{})
	fmt.Println(editor.Type("Привет, мир!"))
}
