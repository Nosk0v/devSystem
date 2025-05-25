-- +goose Up
CREATE TABLE IF NOT EXISTS "goose_db_version" (
                                                  "id" SERIAL PRIMARY KEY,
                                                  "version_id" BIGINT NOT NULL,
                                                  "is_applied" BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "Organization" (
                                              "organization_id" SERIAL PRIMARY KEY,
                                              "name" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "MaterialType" (
                                              "type_id" SERIAL PRIMARY KEY,
                                              "type" VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Competency" (
                                            "competency_id" SERIAL PRIMARY KEY,
                                            "name" VARCHAR(255) NOT NULL UNIQUE,
                                            "description" TEXT NOT NULL,
                                            "parent_id" INTEGER REFERENCES "Competency"("competency_id") ON DELETE SET NULL,
                                            "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "Material" (
                                          "material_id" SERIAL PRIMARY KEY,
                                          "title" VARCHAR(255) NOT NULL,
                                          "description" TEXT NOT NULL,
                                          "type" INTEGER REFERENCES "MaterialType"("type_id") ON DELETE SET NULL,
                                          "content" TEXT NOT NULL,
                                          "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "MaterialCompetency" (
                                                    "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                    "competency_id" INTEGER NOT NULL REFERENCES "Competency"("competency_id") ON DELETE CASCADE,
                                                    PRIMARY KEY ("material_id", "competency_id")
);



CREATE TABLE IF NOT EXISTS "Role" (
                                      "id" SERIAL PRIMARY KEY,
                                      "name" VARCHAR
);

CREATE TABLE IF NOT EXISTS "Account" (
                                         "email" VARCHAR PRIMARY KEY,
                                         "password" VARCHAR,
                                         "name" VARCHAR,
                                         "role" INTEGER NOT NULL REFERENCES "Role"("id"),
                                         "organization_id" INTEGER REFERENCES "Organization"("organization_id") ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS "Course" (
                                        "course_id" SERIAL PRIMARY KEY,
                                        "title" VARCHAR(255) NOT NULL UNIQUE,
                                        "description" TEXT NOT NULL,
                                        "create_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        "created_by" VARCHAR REFERENCES "Account"("email") ON DELETE SET NULL,
                                        "organization_id" INTEGER REFERENCES "Organization"("organization_id") ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS "CourseMaterial" (
                                                "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                PRIMARY KEY ("course_id", "material_id")
);

CREATE TABLE IF NOT EXISTS "CourseCompetency" (
                                                  "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                  "competency_id" INTEGER NOT NULL REFERENCES "Competency"("competency_id") ON DELETE CASCADE,
                                                  PRIMARY KEY ("course_id", "competency_id")
);


INSERT INTO "Organization" ("name") VALUES
                                        ('ООО Альфа'),
                                        ('ЗАО Омега');


INSERT INTO "MaterialType" ("type")
VALUES
    ('Статья'),
    ('Книга'),
    ('Видео');

INSERT INTO "Material" ("title", "description", "type", "content") VALUES
                                                                       ('Введение в разработку ПО', 'Основы подходов к разработке программ', 2, 'Содержимое книги о разработке'),
                                                                       ('Разработка Backend-сервисов', 'Принципы создания серверных решений', 1, 'Статья по backend'),
                                                                       ('Frontend: первые шаги', 'Базовые принципы UI-разработки', 1, 'Материал о frontend'),
                                                                       ('Проектирование БД', 'Как спроектировать эффективную БД', 2, 'Глава из книги по БД'),
                                                                       ('PostgreSQL на практике', 'Работа с PostgreSQL шаг за шагом', 3, 'Видео-практикум PostgreSQL'),
                                                                       ('React для начинающих', 'Интерактивный курс по React', 3, 'Видеоуроки по React'),
                                                                       ('Проектирование REST API', 'Создание удобных API', 1, 'Статья про API'),
                                                                       ('Безопасность в разработке', 'Базовые методы защиты данных', 1, 'Инструкции по безопасности'),
                                                                       ('Git на практике', 'Работа с ветками и коммитами', 2, 'Книга о Git'),
                                                                       ('CI/CD и DevOps', 'Настройка автоматизации', 3, 'Видеогайд по DevOps'),
                                                                       ('Тестирование кода', 'Как писать unit-тесты', 1, 'Пример статьи с тестами'),
                                                                       ('Дизайн интерфейсов', 'UX и UI в работе разработчика', 2, 'Учебное пособие по дизайну'),
                                                                       ('Адаптивная верстка', 'Создание responsive сайтов', 3, 'Видеоурок'),
                                                                       ('Typescript в проектах', 'Как перейти на TS', 1, 'Статья по TS'),
                                                                       ('Сервер на Node.js', 'Обзор возможностей Node.js', 2, 'Книга'),
                                                                       ('Контейнеризация Docker', 'Запуск приложений в контейнерах', 3, 'Видеообзор Docker'),
                                                                       ('Гибкие методологии', 'Scrum и Kanban для команд', 1, 'Методичка по agile'),
                                                                       ('Интеграция API', 'Связь между сервисами через API', 1, 'Статья с примерами'),
                                                                       ('Оптимизация производительности', 'Как ускорить веб-приложение', 2, 'Руководство по оптимизации'),
                                                                       ('Аутентификация и JWT', 'Реализация входа по токенам', 3, 'Видеоинструкция');

INSERT INTO "Competency" ("name", "description", "parent_id") VALUES
                                                                  ('Software Development', 'Основы проектирования и реализации программных решений', NULL),
                                                                  ('Backend Engineering', 'Разработка серверной логики и API-интерфейсов', 1),
                                                                  ('Frontend Engineering', 'Создание пользовательских интерфейсов веб-приложений', 1),
                                                                  ('Database Design', 'Проектирование и нормализация структур баз данных', 1),
                                                                  ('PostgreSQL', 'Работа с реляционной СУБД PostgreSQL', 4),
                                                                  ('React.js', 'Разработка SPA с использованием React', 3),
                                                                  ('REST API Design', 'Создание RESTful API с учетом лучших практик', 2),
                                                                  ('Security Basics', 'Основы безопасности приложений и защита данных', 2),
                                                                  ('Version Control', 'Использование Git и систем контроля версий', 1),
                                                                  ('DevOps Fundamentals', 'Настройка CI/CD и автоматизация процессов', 2),
                                                                  ('Unit Testing', 'Написание и запуск модульных тестов', 2),
                                                                  ('UX/UI Principles', 'Базовые принципы удобства и дизайна интерфейсов', 3),
                                                                  ('Responsive Design', 'Адаптивная вёрстка под разные устройства', 3),
                                                                  ('Typescript', 'Использование строгой типизации в разработке', 3),
                                                                  ('Node.js', 'Серверная разработка на Node.js', 2),
                                                                  ('Docker Basics', 'Контейнеризация приложений с помощью Docker', 10),
                                                                  ('Agile Practices', 'Гибкие методологии разработки, Scrum, Kanban', NULL),
                                                                  ('API Integration', 'Интеграция сторонних API в приложения', 7),
                                                                  ('Performance Optimization', 'Повышение производительности веб-приложений', 2),
                                                                  ('Authentication', 'Реализация регистрации и входа с токенами', 8);

INSERT INTO "MaterialCompetency" ("material_id", "competency_id") VALUES
                                                                      (1, 1), (1, 9), (1, 17),
                                                                      (2, 2), (2, 7), (2, 8),
                                                                      (3, 3), (3, 6), (3, 13),
                                                                      (4, 4), (4, 1), (4, 17),
                                                                      (5, 5), (5, 4), (5, 19),
                                                                      (6, 6), (6, 14), (6, 11),
                                                                      (7, 7), (7, 2), (7, 18),
                                                                      (8, 8), (8, 20), (8, 7),
                                                                      (9, 9), (9, 1), (9, 17),
                                                                      (10, 10), (10, 16), (10, 5),
                                                                      (11, 11), (11, 2), (11, 20),
                                                                      (12, 12), (12, 3), (12, 13),
                                                                      (13, 13), (13, 6), (13, 3),
                                                                      (14, 14), (14, 3), (14, 11),
                                                                      (15, 15), (15, 2), (15, 7),
                                                                      (16, 16), (16, 10), (16, 5),
                                                                      (17, 17), (17, 1), (17, 18),
                                                                      (18, 18), (18, 7), (18, 2),
                                                                      (19, 19), (19, 3), (19, 14),
                                                                      (20, 20), (20, 8), (20, 18);
-- Роли
INSERT INTO "Role" ("id", "name")
VALUES
    (1, 'Пользователь'),
    (0, 'Администратор'),
    (2, 'Супер админ');

-- Аккаунты
INSERT INTO "Account" ("email", "password", "name", "role", "organization_id") VALUES
                                                                                   ('1234@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Иванов Иван Иванович', 0, 1),
                                                                                   ('5678@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Петров Петр Петрович', 1, 2),
                                                                                   ('root@system.dev', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Суперадмин', 2, NULL);


INSERT INTO "Course" ("title", "description", "created_by", "organization_id") VALUES
                                                                                   ('Курс по Backend', 'Освоение серверной разработки на Go и PostgreSQL', '1234@mail.ru', 1),
                                                                                   ('Frontend для начинающих', 'Изучение React и вёрстки с нуля', '5678@mail.ru', 2);

-- Связи курс-материал
INSERT INTO "CourseMaterial" ("course_id", "material_id") VALUES
                                                              (1, 2), -- Backend курс → материал о Backend
                                                              (1, 5), -- Backend курс → PostgreSQL
                                                              (2, 3), -- Frontend курс → начало React
                                                              (2, 6); -- Frontend курс → React интерактив

-- Связи курс-компетенция
INSERT INTO "CourseCompetency" ("course_id", "competency_id") VALUES
                                                                  (1, 2), -- Backend Engineering
                                                                  (1, 5), -- PostgreSQL
                                                                  (2, 3), -- Frontend Engineering
                                                                  (2, 6); -- React.js


-- +goose Down
DROP TABLE IF EXISTS "MaterialCompetency";
DROP TABLE IF EXISTS  "Course";
DROP TABLE IF EXISTS  "CourseMaterial";
DROP TABLE IF EXISTS  "CourseCompetency";
DROP TABLE IF EXISTS "Material";
DROP TABLE IF EXISTS "Competency";
DROP TABLE IF EXISTS "MaterialType";
DROP TABLE IF EXISTS "Account";
DROP TABLE IF EXISTS "Role";
DROP TABLE IF EXISTS "goose_db_version";