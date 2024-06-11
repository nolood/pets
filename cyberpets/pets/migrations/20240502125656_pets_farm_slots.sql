-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY UNIQUE,
    tg_id INTEGER NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL,
    lastname VARCHAR(100),
    firstname VARCHAR(100),
    language_code VARCHAR(5),
    is_premium BOOLEAN,
    photo_url VARCHAR(1000),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS farms (
    id SERIAL PRIMARY KEY UNIQUE,
    user_id INTEGER NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS pets (
                                    id SERIAL PRIMARY KEY UNIQUE,
                                    user_id INTEGER,
                                    rarity INTEGER,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS slots (
    id SERIAL PRIMARY KEY UNIQUE,
    farm_id INTEGER NOT NULL,
    pet_id INTEGER,
    charge INTEGER NOT NULL DEFAULT 100,
    index INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (farm_id) REFERENCES farms (id),
    FOREIGN KEY (pet_id) REFERENCES pets (id)
);



CREATE TABLE IF NOT EXISTS eggs (
  id SERIAL PRIMARY KEY UNIQUE,
  rarity INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS farms CASCADE;
DROP TABLE IF EXISTS slots CASCADE;
DROP TABLE IF EXISTS pets CASCADE;
DROP TABLE IF EXISTS eggs CASCADE;
-- +goose StatementEnd
