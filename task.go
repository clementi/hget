package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func TaskPrint() error {
	entries, err := os.ReadDir(filepath.Join(os.Getenv("HOME"), dataFolder))
	if err != nil {
		return err
	}

	taskNames := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			taskNames = append(taskNames, entry.Name())
		}
	}

	if len(taskNames) == 0 {
		log.Println("No tasks")
	} else {
		log.Println("Tasks:")
		fmt.Println(strings.Join(taskNames, "\n"))
	}

	return nil
}

func Resume(task string) (*State, error) {
	return Read(task)
}
