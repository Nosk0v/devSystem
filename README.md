# devSystem API

**devSystem** — это серверная часть системы для управления развитием сотрудников под названием **«Компетентум»**. Этот проект предоставляет API для работы с базой данных, содержащей материалы и компетенции.

---

## 📋 Содержание
- [Требования](#требования)
- [Запуск](#запуск)
  - [Локальный запуск](#локальный-запуск)
  - [Запуск с Docker](#запуск-с-docker)
- [API Документация](#api-документация)
- [Разработчики](#разработчики)

---

## 📦 Требования

### Для локального запуска:
- **Go** версии 1.22 или выше.
- **PostgreSQL** версии 12+.
- **Git** для клонирования репозитория.

### Для запуска с Docker:
- **Docker** версии 20.10+.
- **Docker Compose** версии 1.29+.

---

## 🚀 Запуск

### Локальный запуск

1. **Клонирование репозитория:**
   ```bash
   git clone https://github.com/Nosk0v/devSystem.git
   cd devSystem
2. **Настройка базы данных:**
   
   Cоздайте базу данных
   
   ```sql
   
   CREATE DATABASE development_system;
   
  Скрипт для создания таблиц находится в db/migrations

3. **Настройка конфига:**

    В файле config/config.json укажите параметры подключения к базе данных: (ниже пример)
    ```json
    {
    "host": "localhost",
    "port": "5432",
    "username": "postgres",
    "password": "1234",
    "dbname": "development_system",
    "sslmode": "disable"
    }
5. **Запуск сервера:**
   ```bash
   SKIP_MIGRATIONS=true go run ./cmd/main.go
6. **Проверка работы:**
   
    После запуска приложение доступно на http://localhost:8080.

### Запуск с Docker

1. **Клонирование репозитория:**
   ```bash
   git clone https://github.com/Nosk0v/devSystem.git
   cd devSystem
2. **Сборка и запуск контейнеров:**
   
   ```bash
   docker-compose up --build
3. **Проверка работы:**
   
    После запуска приложение доступно на http://localhost:8080.
   
### 📖 API Документация
API документация сгенерирована с использованием Swagger. Доступна в папке cmd/docs

### 👨‍💻 Разработчики

Александр Носков
- Email: alexandernoskov.dev@gmail.com
- Telegram: @Noskov_dev

   
