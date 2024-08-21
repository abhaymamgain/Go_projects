package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadLine() ([]string, error) {
	file, err := os.Open("prices/prices.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("error reading file:", err)
		file.Close()
		return nil, err
	}
	return lines, nil
}

func WriteFile(data interface{}, path string) {

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		file.Close()
		return
	}

}
