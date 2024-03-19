package note

import (
	"encoding/json"
	"os"
	"fmt"
	"time"
	"errors"
	"strings"
)

type Note struct {
	Title 		string 		`json:"title"`
	Content 	string 		`json:"content"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)

	return nil
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

