package storage

import (
	"encoding/json"
	"os"
	"todo-app/internal/todo/model"
)

func LoadJSON(filename string) ([]model.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveJSON(filename string, tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
