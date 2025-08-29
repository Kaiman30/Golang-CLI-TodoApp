# To-Do менджер на Go

Консольное приложение для управления задачами с поддержкой JSON/CSV и CLI

## Установка

1. Убедитесь, что установлен Go (Версия 1.24+)
2. Клонируйте или распакуйте проект
3. Соберите приложение:

```bash
go build -o todo ./cmd/todo/main.go
```

## Команды

```bash
# Добавить задачу
./todo add --desc="Изучить Go"

# Показать задачи
./todo list --filter=all      # Все
./todo list --filter=pending  # Невыполненные
./todo list --filter=done     # Выполненные

# Отметить задачу выполненной
./todo complete --id=1

# Удалить заадачу
./todo delete --id=1

# Экспорт в JSON/CSV
./todo export --format=json --out=tasks.json
./todo export --format=csv --out=tasks.csv

# Импорт из файла
./todo load --file=tasks.json
./todo load --file=tasks.csv
```

## Хранение данных
По умолчанию данные хранятся в tasks.json, который создается автоматически при первом запуске приложения
