package main

import (
	"bufio"
	"example.com/note/note"
	"example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	saveData(userNote)
	if err != nil {
		return
	}
}

// functions:
func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return err
	}
	fmt.Println("Note saved successfully!")
	return nil

}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
