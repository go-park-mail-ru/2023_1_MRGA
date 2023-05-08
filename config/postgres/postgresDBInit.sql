CREATE ROLE mrga WITH LOGIN PASSWORD 'mrga_password';

-- Переименовать существующую базу данных
ALTER DATABASE mrga RENAME TO mrgaDB;

-- Изменить владельца базы данных на нового пользователя
ALTER DATABASE mrgaDB OWNER TO <new_user>;