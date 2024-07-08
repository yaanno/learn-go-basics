package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(Text string) (Todo, error) {
	if Text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text,
	}, nil
}

func (todo *Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
