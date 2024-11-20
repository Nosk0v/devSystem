-- +goose Up
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS MaterialType
(
    type_id SERIAL PRIMARY KEY,
    type    VARCHAR(50) NOT NULL
);
-- Вставка по-умолчанию в таблицу MaterialType
INSERT INTO MaterialType (type) VALUES
    ('Курс'),
    ('Книга'),
    ('Видео');


CREATE TABLE IF NOT EXISTS Competency
(
    competency_id SERIAL PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    description   TEXT,
    parent_id     INTEGER REFERENCES competency ON DELETE SET NULL,
    create_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Material
(
    material_id SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    type        INTEGER REFERENCES MaterialType ON DELETE SET NULL,
    content     TEXT,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS MaterialCompetency
(
    material_id   INTEGER NOT NULL REFERENCES material ON DELETE CASCADE,
    competency_id INTEGER NOT NULL REFERENCES competency ON DELETE CASCADE,
    PRIMARY KEY (material_id, competency_id)
);

CREATE TABLE IF NOT EXISTS "User"
(
    username VARCHAR(255) NOT NULL PRIMARY KEY
);

-- +goose Down
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS MaterialCompetency;
DROP TABLE IF EXISTS Material;
DROP TABLE IF EXISTS Competency;
DROP TABLE IF EXISTS MaterialType;
DROP TABLE IF EXISTS goose_db_version;