# TODO-лист на Go + Fiber + PostgreSQL

REST API для управления задачами (TODO-лист), написанный на языке Go с использованием веб-фреймворка [Fiber](https://github.com/gofiber/fiber) и базой данных PostgreSQL, развернутой в Docker.

## Стек технологий

- Go 1.24
- Fiber v5
- PostgreSQL (Docker-контейнер)
- YAML для конфигурации
- Поддержка `.env` для задания пути к конфигу

---

## Запуск приложения

### 1. Клонировать репозиторий

```bash
git clone https://github.com/twobakov/todo-service.git
cd todo-service
```

### 2. Настроить .env файл
Создайте файл .env в корне проекта с содержимым:
```env
CONFIG_PATH=config/config.yaml
```
Делается это с помощью команды:
```bash
echo "CONFIG_PATH=config/config.yaml" > .env
```

### 3. Запустить контейнер с базой данных PostgreSQL

```bash
docker run --name postgres_alt \
  -e POSTGRES_PASSWORD=0897 \
  -e POSTGRES_DB=todolist \
  -p 5433:5432 \
  -d postgres
```
### 4. Создать таблицу с заданиями в базе данных:

Подключиться к базе данных в интерактивном режиме:

```bash
docker exec -it postgres_alt psql -U postgres -d todolist
```
Создать таблицу:
```bash
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
```

### 5. Запустить приложение
```bash
go run cmd/api/main.go
```
---

Приложение доступно по адресу: http://localhost:8080

Можно тестировать с помощью:

- Postman

- curl

##  Возможности
- Добавление задач

- Удаление задач

- Обновление статуса

- Получение списка задач

- Статусы: new, in_progress, done
