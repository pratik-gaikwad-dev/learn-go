package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Todo struct {
	Id   uuid.UUID `json:"todo_id"`
	Text string    `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("Note content should not be empty.")
	}
	return &Todo{
		Id:   uuid.New(),
		Text: text,
	}, nil
}

func (n *Todo) Print() {
	fmt.Println("ID: ", n.Id)
	fmt.Println("Todo: ", n.Text)
}

func (n *Todo) PrintAll() error {
	fileName := "todo.json"

	_, err := os.Stat(fileName)

	if err != nil {
		return err
	}

	readFile, readErr := os.ReadFile(fileName)

	if readErr != nil {
		return readErr
	}

	var todos []Todo

	unmErr := json.Unmarshal(readFile, &todos)

	if unmErr != nil {
		return err
	}
	for i := 1; i < len(todos)+1; i++ {
		fmt.Println("Todo: ", i)
		fmt.Println("\t ID: ", todos[i-1].Id)
		fmt.Println("\t Todo: ", todos[i-1].Text)
	}
	return nil
}

func (n *Todo) Save() error {
	fileName := "todo.json"

	_, err := os.Stat(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			crt, crtErr := os.Create(fileName)
			if crtErr != nil {
				return crtErr
			}
			crt.WriteString("[]")
			crt.Close()
		} else {
			return err
		}
	}

	readFile, readErr := os.ReadFile(fileName)

	if readErr != nil {
		return readErr
	}

	var todos []Todo

	unmErr := json.Unmarshal(readFile, &todos)

	if unmErr != nil {
		return err
	}
	todos = append(todos, *n)

	jsonData, marshalErr := json.MarshalIndent(todos, "", "  ")

	if marshalErr != nil {
		return marshalErr
	}

	writeErr := os.WriteFile(fileName, jsonData, 0644)

	if writeErr != nil {
		return writeErr
	}

	return nil
}
