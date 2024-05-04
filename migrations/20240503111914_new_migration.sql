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

INSERT INTO users (tg_id, username) VALUES (1, 'test');

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
                                    image VARCHAR(1000),
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO pets (rarity, image) VALUES (1, '/pet1.png');
INSERT INTO pets (rarity, image) VALUES (2, '/pet2.png');
INSERT INTO pets (rarity, image) VALUES (3, '/pet3.png');
INSERT INTO pets (rarity, image) VALUES (4, '/pet4.png');
INSERT INTO pets (rarity, image) VALUES (5, '/pet5.png');
INSERT INTO pets (rarity, image) VALUES (1, '/pet6.png');
INSERT INTO pets (rarity, image) VALUES (2, '/pet7.png');
INSERT INTO pets (rarity, image) VALUES (3, '/pet8.jpg');
INSERT INTO pets (rarity, image) VALUES (4, '/pet9.jpg');
INSERT INTO pets (rarity, image) VALUES (5, '/pet10.jpg');
INSERT INTO pets (rarity, image) VALUES (1, '/pet11.jpg');
INSERT INTO pets (rarity, image) VALUES (2, '/pet12.jpg');
INSERT INTO pets (rarity, image) VALUES (3, '/pet13.jpg');
INSERT INTO pets (rarity, image) VALUES (4, '/pet14.jpg');
INSERT INTO pets (rarity, image) VALUES (5, '/pet15.jpg');

CREATE TABLE IF NOT EXISTS UsersPets (
                                        id SERIAL PRIMARY KEY UNIQUE,
                                        user_id INTEGER REFERENCES users (id),
                                        pet_id INTEGER REFERENCES pets (id),
                                        level INTEGER NOT NULL DEFAULT 1,

                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
                                    image VARCHAR(1000),
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO eggs (rarity, image) VALUES (1, '/egg1.png');

CREATE TABLE IF NOT EXISTS UsersEggs (
                                        id SERIAL PRIMARY KEY UNIQUE,
                                        user_id INTEGER REFERENCES users (id),
                                        egg_id INTEGER REFERENCES eggs (id),
                                        hatch_time INTEGER NOT NULL,
                                        hatch_start TIMESTAMP,
                                        hatch_end TIMESTAMP,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS incubators (
                                    id SERIAL PRIMARY KEY UNIQUE,
                                    user_id INTEGER REFERENCES users (id),
                                    egg_id INTEGER REFERENCES UsersEggs (id) DEFAULT NULL,
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
DROP TABLE IF EXISTS UsersPets CASCADE;
DROP TABLE IF EXISTS incubators CASCADE;
DROP TABLE IF EXISTS UsersEggs CASCADE;
-- +goose StatementEnd
