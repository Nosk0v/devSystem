-- +goose Up
CREATE TABLE IF NOT EXISTS "goose_db_version" (
                                                  "id" SERIAL PRIMARY KEY,
                                                  "version_id" BIGINT NOT NULL,
                                                  "is_applied" BOOLEAN DEFAULT FALSE
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
                                         "role" INTEGER NOT NULL REFERENCES "Role"("id")
);

INSERT INTO "MaterialType" ("type")
VALUES
    ('Статья'),
    ('Книга'),
    ('Видео');

INSERT INTO "Role" ("id", "name")
VALUES
    (0, 'Пользователь'),
    (1, 'Администратор');

INSERT INTO "Account" ("email", "password", "name", "role")
VALUES
    ('user@test.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Иванов Иван Иванович', 0),
    ('admin@test.ru', '$2a$10$hitarfnbzlubZuZtQKITq.6zoul4yywj1f6Sn0dl.N41uuRwGhXKm', 'Петров Петр Петрович', 1);

-- +goose Down
DROP TABLE IF EXISTS "MaterialCompetency";
DROP TABLE IF EXISTS "Material";
DROP TABLE IF EXISTS "Competency";
DROP TABLE IF EXISTS "MaterialType";
DROP TABLE IF EXISTS "Account";
DROP TABLE IF EXISTS "Role";
DROP TABLE IF EXISTS "goose_db_version";