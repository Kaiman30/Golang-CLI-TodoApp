// internal/todo/manager.go
package todo

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"todo-app/internal/todo/storage"
)

type Manager struct {
	Tasks  []storage.Task
	NextID int
	file   string
	mu     sync.Mutex
}

func NewManager(file string) *Manager {
	manager := &Manager{
		file:   file,
		NextID: 1,
	}
	tasks, err := storage.LoadJSON(file)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Предупреждение: не удалось загрузить задачи:", err)
		}
	} else {
		manager.Tasks = tasks
		if len(tasks) > 0 {
			maxID := 0
			for _, t := range tasks {
				if t.ID > maxID {
					maxID = t.ID
				}
			}
			manager.NextID = maxID + 1
		}
	}
	return manager
}

func (m *Manager) Save() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return storage.SaveJSON(m.file, m.Tasks)
}

func (m *Manager) Add(desc string) error {
	if desc == "" {
		return errors.New("описание задачи не может быть пустым")
	}
	task := storage.Task{
		ID:          m.NextID,
		Description: desc,
		Done:        false,
	}
	m.Tasks = append(m.Tasks, task)
	m.NextID++
	return m.Save()
}

func (m *Manager) List(filter string) ([]storage.Task, error) {
	var result []storage.Task
	switch filter {
	case "all":
		result = m.Tasks
	case "done":
		for _, t := range m.Tasks {
			if t.Done {
				result = append(result, t)
			}
		}
	case "pending":
		for _, t := range m.Tasks {
			if !t.Done {
				result = append(result, t)
			}
		}
	default:
		return nil, fmt.Errorf("неизвестный фильтр: %s", filter)
	}
	return result, nil
}

func (m *Manager) Complete(id int) error {
	for i := range m.Tasks {
		if m.Tasks[i].ID == id {
			m.Tasks[i].Done = true
			return m.Save()
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}

func (m *Manager) Delete(id int) error {
	for i, t := range m.Tasks {
		if t.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return m.Save()
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}
