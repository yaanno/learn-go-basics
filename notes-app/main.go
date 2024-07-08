package main

import (
	"bufio"
	"fmt"
	"notes-app/note"
	"os"
	"strings"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	note.DisplayNote()
	err = note.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Note saved")
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
