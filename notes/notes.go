package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id      uuid.UUID `json:"note_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"created_at"`
}

func New(title string, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("Note title should not be empty.")
	}
	if content == "" {
		return nil, errors.New("Note content should not be empty.")
	}
	return &Note{
		Id:      uuid.New(),
		Title:   title,
		Content: content,
		Time:    time.Now(),
	}, nil
}

func (n *Note) Print() {
	fmt.Println("ID: ", n.Id)
	fmt.Println("Title: ", n.Title)
	fmt.Println("Content: ", n.Content)
	fmt.Println("Time: ", n.Time)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.MarshalIndent(n, "", "  ")

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(fileName, jsonData, 0644)

	return writeErr
}
