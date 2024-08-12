package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the directory path to process: ")
	dirPath, _ := reader.ReadString('\n')
	dirPath = strings.TrimSpace(dirPath)

	err := processDirectory(dirPath)
	if err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		return
	}

	fmt.Println("File names processed successfully!")
}

func processDirectory(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		oldPath := filepath.Join(dirPath, file.Name())
		newName := strings.ReplaceAll(file.Name(), " ", "_")
		newPath := filepath.Join(dirPath, newName)

		if oldPath != newPath {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return fmt.Errorf("error renaming file %s: %v", oldPath, err)
			}
			fmt.Printf("Renamed: %s -> %s\n", oldPath, newPath)
		}
	}

	return nil
}
