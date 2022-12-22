package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var dataFolder = ".hget/"
var stateFileName = "state.json"

type State struct {
	Url   string
	Parts []Part
}

type Part struct {
	Url       string
	Path      string
	RangeFrom int64
	RangeTo   int64
}

func (s *State) Save() error {
	// make temp folder
	// only working in unix with env HOME
	folder := FolderOf(s.Url)
	log.Printf("Saving current download data in %s\n", folder)
	if err := MkdirIfNotExist(folder); err != nil {
		return err
	}

	// move current downloading file to data folder
	for _, part := range s.Parts {
		os.Rename(part.Path, filepath.Join(folder, filepath.Base(part.Path)))
	}

	// save state file
	j, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(folder, stateFileName), j, 0644)
}

func Read(task string) (*State, error) {
	file := filepath.Join(os.Getenv("HOME"), dataFolder, task, stateFileName)
	log.Printf("Getting data from %s\n", file)
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	s := new(State)
	err = json.Unmarshal(bytes, s)
	return s, err
}

func Delete(task string) error {
	dataPath := filepath.Join(os.Getenv("HOME"), dataFolder)
	if err := MkdirIfNotExist(dataPath); err != nil {
		return err
	}
	taskPath := filepath.Join(dataPath, task)
	if DirExists(taskPath) {
		log.Printf("Deleting task %s\n", taskPath)
		return os.RemoveAll(taskPath)
	} else {
		return fmt.Errorf("task '%s' does not exist", task)
	}
}
