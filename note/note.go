package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"elif.com/fileoperations"
)

type Note struct {
    Title     string    `json:"title"`      // converted as "title" in JSON
    Content   string    `json:"content"`    
    CreatedAt time.Time `json:"created_at"`
}

func New (title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note {}, errors.New("Title or content cannot be empty")

	}
	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil

}

func (note Note) DisplayContent () {
	fmt.Printf("Title:\n %v \n\n Content: \n %v \n\n Created at: \n %v \n\n", note.Title, note.Content, note.CreatedAt)
}

func (note Note) SaveToFile () error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	fileoperations.WriteToFile(json, fileName, 0644)
	return nil
}
