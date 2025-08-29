package main

import (
	"flag"
	"fmt"
	"os"

	"todo-app/internal/todo"
	"todo-app/internal/todo/storage"
)

func main() {
	// Определение команд
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	loadCmd := flag.NewFlagSet("load", flag.ExitOnError)

	// Флаги для команд
	desc := addCmd.String("desc", "", "Описание задачи")
	filter := listCmd.String("filter", "all", "Фильтр: all/done/pending")
	id := completeCmd.Int("id", 0, "ID задачи для завершения")
	delID := deleteCmd.Int("id", 0, "ID задачи для удаления")
	format := exportCmd.String("format", "json", "Формат экспорта: json/csv")
	out := exportCmd.String("out", "tasks.json", "Путь для экспорта")
	file := loadCmd.String("file", "", "Путь к файлу для импорта")

	// Проверка переданной команды
	if len(os.Args) < 2 {
		fmt.Println("Укажите команду: add, list, complete, delete, export, load")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *desc == "" {
			fmt.Println("Флаг --desc обязателен")
			os.Exit(1)
		}
		manager := todo.NewManager("tasks.json")
		if err := manager.Add(*desc); err != nil {
			fmt.Println("Ошибка добавления задачи:", err)
			os.Exit(1)
		}
		fmt.Println("Задача добавлена")

	case "list":
		listCmd.Parse(os.Args[2:])
		manager := todo.NewManager("tasks.json")
		tasks, err := manager.List(*filter)
		if err != nil {
			fmt.Println("Ошибка при получении списка:", err)
			os.Exit(1)
		}
		for _, t := range tasks {
			status := "pending"
			if t.Done {
				status = "done"
			}
			fmt.Printf("[%d] %s [%s]\n", t.ID, t.Description, status)
		}

	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *id == 0 {
			fmt.Println("Флаг --id обязателен")
			os.Exit(1)
		}
		manager := todo.NewManager("tasks.json")
		if err := manager.Complete(*id); err != nil {
			fmt.Println("Ошибка при завершении задачи:", err)
			os.Exit(1)
		}
		fmt.Println("Задача отмечена как выполненная")

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *delID == 0 {
			fmt.Println("Флаг --id обязателен")
			os.Exit(1)
		}
		manager := todo.NewManager("tasks.json")
		if err := manager.Delete(*delID); err != nil {
			fmt.Println("Ошибка при удалении задачи:", err)
			os.Exit(1)
		}
		fmt.Println("Задача удалена")

	case "export":
		exportCmd.Parse(os.Args[2:])
		manager := todo.NewManager("tasks.json")
		var err error
		switch *format {
		case "json":
			err = storage.SaveJSON(*out, manager.Tasks)
		case "csv":
			err = storage.SaveCSV(*out, manager.Tasks)
		default:
			fmt.Println("Неподдерживаемый формат: --format=json/csv")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Ошибка экспорта:", err)
			os.Exit(1)
		}
		fmt.Printf("Данные экспортированы в %s\n", *out)

	case "load":
		loadCmd.Parse(os.Args[2:])
		if *file == "" {
			fmt.Println("Флаг --file обязателен")
			os.Exit(1)
		}
		manager := todo.NewManager("tasks.json")
		var err error
		ext := *file
		if len(ext) > 4 && ext[len(ext)-4:] == ".csv" {
			manager.Tasks, err = storage.LoadCSV(*file)
		} else {
			manager.Tasks, err = storage.LoadJSON(*file)
		}
		if err != nil {
			fmt.Println("Ошибка импорта:", err)
			os.Exit(1)
		}
		if err := manager.Save(); err != nil {
			fmt.Println("Ошибка сохранения после загрузки:", err)
			os.Exit(1)
		}
		fmt.Println("Данные загружены")

	default:
		fmt.Println("Неизвестная команда. Доступные: add, list, complete, delete, export, load")
		os.Exit(1)
	}
}
