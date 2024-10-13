package todo

import (
	"encoding/json"
	"errors"
	"fmt"

	"elif.com/fileoperations"
)

type Todo struct {
    Text   string    `json:"text"`    
}

func New (text string) (Todo, error) {
	if text == "" {
		return Todo {}, errors.New("Text cannot be empty")

	}
	return Todo {
		Text: text,
	}, nil

}

func (todo Todo) DisplayContent () {
	fmt.Println(todo.Text)
}

func (todo Todo) SaveToFile () error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	fileoperations.WriteToFile(json, fileName, 0644)
	return nil
}
