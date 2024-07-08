package main

import (
	"bufio"
	"fmt"
	"notes-app/note"
	"notes-app/todo"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(&todo)
	if err != nil {
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	err = outputData(&note)
	if err != nil {
		return
	}
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println(intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println(floatVal)
		return
	}

	switch value.(type) {
	case int:
		fmt.Println(value)
	default:
		break
	}
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Save succeeded")
	return nil
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func getTodoData() string {
	text := getUserInput("Todo: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
