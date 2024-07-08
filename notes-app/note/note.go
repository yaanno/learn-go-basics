package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display() {
	fmt.Printf("Note Title:\n %v\n", note.Title)
	fmt.Printf("Note Content:\n %v\n", note.Content)
	fmt.Println(note.CreatedAt.Local())
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
