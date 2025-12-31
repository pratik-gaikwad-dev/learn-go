package main

import (
	"fmt"

	"example.com/note/notes"
	"example.com/note/todo"
	"example.com/note/utils"
)

func main() {
	// Print Generic -- Can pass both values of same type only not different type
	// utils.PrintGeneric(1, 1) // Ok
	// utils.PrintGeneric(1, "str") // Not Ok

	// Print Any
	// utils.PrintValue(1)
	// utils.PrintValue(1.1)
	// utils.PrintValue("Anything")

	// Todo
	text := utils.GetInput("Enter Task: ")

	// Notes
	title := utils.GetInput("Enter note title: ")
	content := utils.GetInput("Enter note content: ")

	newNote, err := notes.New(title, content)
	newTodo, todoErr := todo.New(text)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if todoErr != nil {
		fmt.Println(todoErr)
	}

	fmt.Println()

	/*
		saveErr := newNote.Save()
		todoSaveErr := newTodo.Save()

		if saveErr != nil {
			fmt.Println(saveErr)
		}
		if todoSaveErr != nil {
			fmt.Println(todoSaveErr)
		}
		fmt.Println("Notes and Todo's saved successfully.")

		Instead of above code the following code is enough, Thanks to interfaces.
	*/
	noteErr := utils.SaveData(newNote)
	todoSaveErr := utils.SaveData(newTodo)
	if noteErr != nil {
		fmt.Println(noteErr)
	}
	if todoSaveErr != nil {
		fmt.Println(todoSaveErr)
	}

	// newTodo.PrintAll()

}
