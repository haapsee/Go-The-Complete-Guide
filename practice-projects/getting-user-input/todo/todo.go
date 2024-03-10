package todo

import (
	"encoding/json"
	"os"
	"fmt"
	"errors"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)

	return nil
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: text,
	}, nil
}

