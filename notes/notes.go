package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

func (n *Note) PrintAll() error {
	fileName := "notes.json"

	_, err := os.Stat(fileName)

	if err != nil {
		return err
	}

	readFile, readErr := os.ReadFile(fileName)

	if readErr != nil {
		return readErr
	}

	var notes []Note

	unmErr := json.Unmarshal(readFile, &notes)

	if unmErr != nil {
		return err
	}
	for i := 1; i < len(notes)+1; i++ {
		fmt.Println("Note: ", i)
		fmt.Println("\t ID: ", notes[i-1].Id)
		fmt.Println("\t Title: ", notes[i-1].Title)
		fmt.Println("\t Content: ", notes[i-1].Content)
		fmt.Println("\t Time: ", notes[i-1].Time)
	}
	return nil
}

func (n *Note) Save() error {
	fileName := "notes.json"

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

	var notes []Note

	unmErr := json.Unmarshal(readFile, &notes)

	if unmErr != nil {
		return err
	}
	notes = append(notes, *n)

	jsonData, marshalErr := json.MarshalIndent(notes, "", "  ")

	if marshalErr != nil {
		return marshalErr
	}

	writeErr := os.WriteFile(fileName, jsonData, 0644)

	if writeErr != nil {
		return writeErr
	}

	return nil
}
