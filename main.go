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

type Outputtable interface {
	Saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	TypedVal, ok := value.(int)

	if ok {
		fmt.Println("integer", TypedVal)
		return
	}
	Floatval, ok := value.(float64)

	if ok {
		fmt.Println("integer", Floatval)
		return
	}
	//case int:
	//	fmt.Println("This is an integer", value)
	//case float64:
	//	fmt.Println("This is an float", value)
	//case string:
	//	fmt.Println("This is an string", value)

	//}
	//fmt.Println(value)
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
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
