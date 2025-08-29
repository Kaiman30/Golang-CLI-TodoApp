package storage

import (
	"encoding/json"
	"os"
)

func LoadJSON(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveJSON(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
