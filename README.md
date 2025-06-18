# devSystem API

**devSystem** — это серверная часть системы для управления развитием сотрудников под названием **«Компетентум»**. Этот проект предоставляет API для работы с базой данных, содержащей материалы и компетенции.

---

## 📋 Содержание
- [Требования](#требования)
- [Запуск](#запуск)
  - [Локальный запуск](#локальный-запуск)
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

### Локальный запуск сервера

1. **Клонирование репозитория:**
   ```bash
   git clone https://github.com/Nosk0v/devSystem.git
   cd devSystem
2. **Настройка базы данных:**
   
   Cоздайте базу данных
   
   ```sql
   
   CREATE DATABASE "development-system-db";
   
  Скрипт для создания таблиц находится в db/migrations

3. **Настройка конфига:**

    В файле config/config.json укажите параметры подключения к базе данных: (ниже пример)
    ```json
    {
      "server": {
      "server_port": "25502",
      "jwt_secret_key": "secret"
      },
      "development-system-database": {
      "db_host": "postgres",
      "db_port": "5432",
      "db_username": "postgres",
      "db_password": "1234",
      "db_name": "development-system-db",
      "db_ssl_mode": "disable"
      }
   }

4. **Запуск сервера локально:**
   ```bash
   cd cmd
   SKIP_MIGRATIONS=true go run main.go
   
5. **Проверка работы:**
   
    После запуска сервер доступен на http://localhost:25502.
   
### 📖 API Документация
API документация сгенерирована с использованием Swagger. Доступна в папке cmd/docs

### 👨‍💻 Разработчики

Александр Носков
- Email: alexandernoskov.dev@gmail.com
- Telegram: @Noskov_dev

   
