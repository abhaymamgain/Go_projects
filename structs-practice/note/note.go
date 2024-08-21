package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")
	}
	return Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}
func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}
