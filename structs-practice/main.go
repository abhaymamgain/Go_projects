package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"struct_practice.com/notes/note"
)

func main() {
	title, content := getNotesData()
	note, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = note.Save()
	if err != nil {
		fmt.Println("note not saved.")
		return
	}
	fmt.Print("Note saved success")
}
func getNotesData() (string, string) {
	title := getUserInput("notes title")
	content := getUserInput("notes content")

	return title, content
}
func getUserInput(prompt string) string {

	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
