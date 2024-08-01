-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL 
-- );
-- CREATE TABLE posts (
--     id SERIAL PRIMARY KEY,
--     users_id 
-- )
-- Создаем таблицу users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Создаем таблицу posts и устанавливаем внешний ключ, связывающий posts с users
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    CONSTRAINT fk_user -- имя ограничения внешнего ключа
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);

INSERT INTO users (id,name) VALUES (1,'IVANOV');