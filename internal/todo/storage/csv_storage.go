package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"todo-app/internal/todo/model"
)

func LoadCSV(filename string) ([]model.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	// Пропускаем заголовок
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) != 3 {
			continue
		}

		id, _ := strconv.Atoi(record[0])
		done, _ := strconv.ParseBool(record[2])
		tasks = append(tasks, model.Task{
			ID:          id,
			Description: record[1],
			Done:        done,
		})
	}

	return tasks, nil
}

func SaveCSV(filename string, tasks []model.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовок
	if err := writer.Write([]string{"id", "description", "done"}); err != nil {
		return err
	}

	for _, t := range tasks {
		doneStr := "false"
		if t.Done {
			doneStr = "true"
		}
		if err := writer.Write([]string{
			strconv.Itoa(t.ID),
			t.Description,
			doneStr,
		}); err != nil {
			return err
		}
	}
	return nil
}
