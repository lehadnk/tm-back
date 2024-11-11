package domain

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileSystemManager struct {
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading file:" + err.Error())
	}

	return data, nil
}

func WriteToFile(file []byte, path string) error {
	err := os.WriteFile(path, file, 0644)
	if err != nil {
		log.Fatalln("Error writing to file: " + err.Error())
	}
	return err
}
