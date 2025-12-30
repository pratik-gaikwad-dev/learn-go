package main

import (
	"fmt"

	"example.com/note/notes"
)

func main() {
	title := notes.GetInput("Enter note title: ")
	content := notes.GetInput("Enter note content: ")

	newNote, err := notes.New(title, content)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	newNote.Print()

	saveErr := newNote.Save()

	if saveErr != nil {
		fmt.Println(saveErr)
	}

	fmt.Println("Note saved successfully.")
}
