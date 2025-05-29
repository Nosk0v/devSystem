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


CREATE TABLE IF NOT EXISTS "MaterialProgress" (
                                                  "user_email" VARCHAR NOT NULL REFERENCES "Account"("email") ON DELETE CASCADE,
                                                  "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                  "material_id" INTEGER NOT NULL REFERENCES "Material"("material_id") ON DELETE CASCADE,
                                                  "is_viewed" BOOLEAN DEFAULT TRUE,
                                                  "viewed_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                  PRIMARY KEY ("user_email", "course_id", "material_id")
);

CREATE TABLE IF NOT EXISTS "CourseProgress" (
                                                "progress_id" SERIAL PRIMARY KEY,
                                                "user_email" VARCHAR NOT NULL REFERENCES "Account"("email") ON DELETE CASCADE,
                                                "course_id" INTEGER NOT NULL REFERENCES "Course"("course_id") ON DELETE CASCADE,
                                                "is_completed" BOOLEAN DEFAULT FALSE,
                                                "completed_at" TIMESTAMP,
                                                UNIQUE("user_email", "course_id")
);

CREATE TABLE "RegistrationPrefix" (
                                      prefix TEXT PRIMARY KEY,                     -- напр. "MTI", "SBER"
                                      organization_id INT REFERENCES "Organization"(organization_id),
                                      created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE "InviteCode" (
                              code TEXT PRIMARY KEY,                        -- напр. "MTI-XYZ123"
                              prefix TEXT REFERENCES "RegistrationPrefix"(prefix),
                              role INT NOT NULL DEFAULT 1,                  -- 1=user, 2=admin
                              used BOOLEAN NOT NULL DEFAULT FALSE,
                              expires_at TIMESTAMP,                         -- NULL = бессрочный
                              created_at TIMESTAMP DEFAULT now()
);

INSERT INTO "Organization" ("name") VALUES
                                        ('ООО Альфа'),
                                        ('ЗАО Омега');


INSERT INTO "MaterialType" ("type")
VALUES
    ('Статья'),
    ('Книга'),
    ('Видео');

INSERT INTO "RegistrationPrefix" ("prefix", "organization_id")
VALUES
    ('ALFA', 1),  -- ООО Альфа
    ('OMEGA', 2); -- ЗАО Омега

INSERT INTO "Material" ("title", "description", "type", "content") VALUES
                                                                        ('Разработка Backend-сервисов', 'Принципы создания серверных решений', 1, 'Backend-разработка требует понимания архитектуры, масштабируемости и безопасности. Один из ключевых аспектов — правильная организация слоёв приложения и взаимодействие с базой данных. Подробнее об этом можно прочитать в статье: https://habr.com/ru/company/otus/blog/539460/'),
                                                                        ('Frontend: первые шаги', 'Базовые принципы UI-разработки', 1, 'Начинающим фронтенд-разработчикам важно освоить основы HTML, CSS и JavaScript. Хорошее введение в тему представлено в статье: https://developer.mozilla.org/ru/docs/Learn/Front-end_web_developer'),
                                                                        ('Проектирование REST API', 'Создание удобных API', 1, 'REST API должны быть интуитивно понятными и хорошо документированными. Рекомендую ознакомиться с руководством по проектированию RESTful API: https://restfulapi.net/'),
                                                                        ('Безопасность в разработке', 'Базовые методы защиты данных', 1, 'Безопасность — неотъемлемая часть разработки. Основы защиты веб-приложений описаны в статье: https://owasp.org/www-project-top-ten/'),
                                                                        ('Тестирование кода', 'Как писать unit-тесты', 1, 'Unit-тестирование помогает выявлять ошибки на ранних этапах разработки. Ознакомьтесь с руководством по написанию тестов на JavaScript: https://jestjs.io/docs/getting-started'),
                                                                        ('Typescript в проектах', 'Как перейти на TS', 1, 'TypeScript добавляет статическую типизацию в JavaScript, что повышает надёжность кода. Подробнее о переходе на TypeScript: https://www.typescriptlang.org/docs/handbook/migrating-from-javascript.html'),
                                                                        ('Гибкие методологии', 'Scrum и Kanban для команд', 1, 'Agile-методологии помогают командам адаптироваться к изменениям. Ознакомьтесь с основами Scrum и Kanban: https://www.atlassian.com/agile'),
                                                                        ('Интеграция API', 'Связь между сервисами через API', 1, 'Интеграция API позволяет различным сервисам взаимодействовать между собой. Руководство по интеграции API: https://www.redhat.com/en/topics/api/what-is-api-integration'),
                                                                        ('Введение в разработку ПО', 'Основы подходов к разработке программ', 2, 'Книга "Алгоритмы. Руководство по разработке" Стивена Скиены предоставляет фундаментальные знания о структуре и разработке программного обеспечения. Подробнее: https://book24.ru/product/algoritmy-rukovodstvo-po-razrabotke-6071303/'),
                                                                        ('Проектирование БД', 'Как спроектировать эффективную БД', 2, 'Эффективное проектирование баз данных требует понимания нормализации, индексации и транзакций. Рекомендуется книга "Проектирование реляционных баз данных" Пола Нильсена.'),
                                                                        ('Git на практике', 'Работа с ветками и коммитами', 2, 'Книга "Pro Git" Скотта Шакона — исчерпывающее руководство по Git, охватывающее все аспекты работы с системой контроля версий. Доступна онлайн: https://git-scm.com/book/ru/v2'),
                                                                        ('Дизайн интерфейсов', 'UX и UI в работе разработчика', 2, 'Книга "Dont Make Me Think" Стива Круга — классика в области UX-дизайна, объясняющая принципы создания удобных интерфейсов.'),
                                                                        ('Сервер на Node.js', 'Обзор возможностей Node.js', 2, 'Книга "Node.js в действии" Майка Кэнтора предоставляет глубокое понимание разработки серверных приложений на Node.js.'),
                                                                        ( 'Оптимизация производительности', 'Как ускорить веб-приложение', 2, 'Оптимизация производительности включает в себя анализ узких мест и применение лучших практик. Рекомендуется книга "High Performance Browser Networking" Ильи Григорика.'),
                                                                        ('PostgreSQL на практике', 'Работа с PostgreSQL шаг за шагом', 3, 'Полный курс по PostgreSQL для начинающих: https://www.youtube.com/watch?v=qw--VYLpxG4'),
                                                                        ('React для начинающих', 'Интерактивный курс по React', 3, 'Видеокурс по React для начинающих: https://www.youtube.com/watch?v=SqcY0GlETPk'),
                                                                        ('CI/CD и DevOps', 'Настройка автоматизации', 3, 'Введение в CI/CD и практики DevOps: https://www.youtube.com/watch?v=1hHMwLxN6EM'),
                                                                        ('Адаптивная верстка', 'Создание responsive сайтов', 3, 'Урок по адаптивной верстке: https://www.youtube.com/watch?v=VQraviuwbzU'),
                                                                        ('Контейнеризация Docker', 'Запуск приложений в контейнерах', 3, 'Обзор Docker и контейнеризации: https://www.youtube.com/watch?v=ZWReEReO2xE'),
                                                                        ('Аутентификация и JWT', 'Реализация входа по токенам', 3, 'Руководство по JWT-аутентификации: https://www.youtube.com/watch?v=7Q17ubqLfaM');


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
                                                                                   ('2org@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Иванов Иван Иванович', 0, 2),
                                                                                   ('5678@mail.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Петров Петр Петрович', 1, 2),
                                                                                   ('root@system.dev', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Суперадмин', 2, NULL);


INSERT INTO "Course" ("title", "description", "created_by", "organization_id") VALUES
                                                                                ('Курс по Backend', 'Освоение серверной разработки на Go и PostgreSQL', '1234@mail.ru', 1),
                                                                                ('Frontend для начинающих', 'Изучение React и вёрстки с нуля', '5678@mail.ru', 2),
                                                                                ('Основы Golang', 'Курс по основам языка программирования Go', '1234@mail.ru', 1),
                                                                                ('Микросервисы на Go', 'Создание микросервисной архитектуры', '1234@mail.ru', 1),
                                                                                ('GRPC и Protocol Buffers', 'Работа с gRPC и Protobuf в бэкенде', '1234@mail.ru', 1),
                                                                                ('JWT и безопасность API', 'Реализация авторизации на JWT', '1234@mail.ru', 1),
                                                                                ('Docker для разработчиков', 'Контейнеризация сервисов и Dockerfile', '1234@mail.ru', 1),
                                                                                ('CI/CD с GitHub Actions', 'Автоматизация сборок и деплоя', '1234@mail.ru', 1),
                                                                                ('SQL и оптимизация запросов', 'Повышение производительности SQL-запросов', '1234@mail.ru', 1),
                                                                                ('PostgreSQL в бою', 'Использование расширений и индексов', '1234@mail.ru', 1),
                                                                                ('Работа с REST API', 'Создание и тестирование REST-интерфейсов', '1234@mail.ru', 1),
                                                                                ('Тестирование бэкенда', 'Unit, интеграционные и e2e тесты', '1234@mail.ru', 1),
                                                                                ('Проектирование БД', 'Связи, нормализация, схема', '1234@mail.ru', 1),
                                                                                ('Введение в DevOps', 'Как связаны разработка и эксплуатация', '1234@mail.ru', 1),
                                                                                ('Мониторинг и алертинг', 'Grafana, Prometheus, alertmanager', '1234@mail.ru', 1),
                                                                                ('Методологии Scrum и Agile', 'Управление командной разработкой', '1234@mail.ru', 1),
                                                                                ('Секреты производительности Go', 'Профилирование и ускорение кода', '1234@mail.ru', 1);

-- Связи курс-материал
INSERT INTO "CourseMaterial" ("course_id", "material_id") VALUES
                                                              (1, 2),
                                                              (1, 5),
                                                              (2, 3),
                                                              (2, 6),
                                                              (3, 1), (3, 2),
                                                              (4, 5), (4, 7),
                                                              (5, 16), (5, 10),
                                                              (6, 10), (6, 11),
                                                              (7, 4), (7, 19),
                                                              (8, 5), (8, 4),
                                                              (9, 7), (9, 18),
                                                              (10, 11), (10, 14),
                                                              (11, 4), (11, 1),
                                                              (12, 10), (12, 16),
                                                              (13, 20), (13, 5),
                                                              (14, 17), (14, 1),
                                                              (15, 2), (15, 19),
                                                              (16, 1), (16, 15),
                                                              (17, 2), (17, 14);

-- Связи курс-компетенция
INSERT INTO "CourseCompetency" ("course_id", "competency_id") VALUES
                                                                  (1, 2),
                                                                  (1, 5),
                                                                  (2, 3),
                                                                  (2, 6),
                                                                  (3, 2), (3, 14),
                                                                  (4, 20), (4, 8),
                                                                  (5, 16), (5, 10),
                                                                  (6, 10), (6, 2),
                                                                  (7, 4), (7, 19),
                                                                  (8, 5), (8, 4),
                                                                  (9, 7), (9, 18),
                                                                  (10, 11), (10, 20),
                                                                  (11, 4), (11, 1),
                                                                  (12, 10), (12, 5),
                                                                  (13, 20), (13, 8),
                                                                  (14, 17), (14, 1),
                                                                  (15, 2), (15, 7),
                                                                  (16, 1), (16, 19),
                                                                  (17, 14), (17, 2);


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
DROP TABLE IF EXISTS "CourseProgress";
DROP TABLE IF EXISTS "goose_db_version";