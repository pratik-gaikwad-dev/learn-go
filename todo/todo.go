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

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Note content should not be empty.")
	}
	return Todo{
		Id:   uuid.New(),
		Text: text,
	}, nil
}

func (n *Todo) Print() {
	fmt.Println("ID: ", n.Id)
	fmt.Println("Content: ", n.Text)
}

func (n *Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.MarshalIndent(n, "", "  ")

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(fileName, jsonData, 0644)

	return writeErr
}
