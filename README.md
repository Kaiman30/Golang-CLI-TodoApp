# To-Do CLI менджер на Go

Консольное приложение для управления задачами с поддержкой JSON/CSV и CLI

---

## Возможности

- Добавление, просмотр, завершение и удаление задач
- Фильтрация по статусу: `все`, `выполненные`, `ожидающие`
- Автоматическое сохранение в `tasks.json`
- Экспорт в форматах **JSON** и **CSV**
- Импорт данных из **JSON** и **CSV**
- Модульная архитектура и обработка ошибок
- Работает в терминале на любой ОС (Windows, macOS, Linux)

---

## Установка

1. Убедитесь, что у вас установлен **Go** (Версия 1.24+)
    ```bash
    go version
    ```
2. Клонируйте или распакуйте проект
    ```bash
    git clone https://github.com/Kaiman30/Golang-CLI-TodoApp.git
    cd todo-app
    ```
3. Соберите приложение:
    ```bash
    go build -o todo ./cmd/todo/main.go
    ```

**Теперь у вас есть бинарник `todo` (или `todo.exe` для Windows)**

---

## Как использовать

```bash
# Добавить задачу
./todo add --desc="Изучить Go"

# Показать задачи
./todo list --filter=all      # Все
./todo list --filter=pending  # Невыполненные
./todo list --filter=done     # Выполненные

# Отметить задачу выполненной
./todo complete --id=1

# Удалить задачу
./todo delete --id=1

# Экспорт в JSON/CSV
./todo export --format=json --out=tasks.json
./todo export --format=csv --out=tasks.csv

# Импорт из файла
./todo load --file=tasks.json
./todo load --file=tasks.csv
```

---

## Хранение данных
По умолчанию данные хранятся в `tasks.json`, который создается автоматически при первом запуске приложения

---

## Пример работы
```bash
# Добавим пару задач
./todo add --desc="Написать README"
./todo add --desc="Собрать проект"

# Посмотрим невыполненные
./todo list --filter=pending

# Отметим первую как выполненную
./todo complete --id=1

# Удалим вторую
./todo delete --id=2

# Экспортируем в CSV
./todo export --format=csv --out=export.csv
```

---

## Требования

- **Go 1.24+**
- Терминал/командная строка
- VS Code (или любой другой редактор кода)