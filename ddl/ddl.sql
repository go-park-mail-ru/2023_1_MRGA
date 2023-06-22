-- cities, educations, hashtags, jobs, reactions, reasons, sexes, statuses, zodiacs -
-- эти таблицы представляют собой справочники и соответствуют 3NF.
-- У них есть первичные ключи, и все их атрибуты зависят от этого ключа.
-- Также не существует транзитивных зависимостей.

CREATE TABLE cities (
                        id bigserial NOT NULL PRIMARY KEY,
                        city text
);

CREATE TABLE complaints (
                            id bigserial NOT NULL UNIQUE PRIMARY KEY,
                            user_id bigint,
                            count bigint
);


CREATE TABLE educations (
                            id bigserial NOT NULL PRIMARY KEY,
                            education text
);


CREATE TABLE hashtags (
                          id bigserial NOT NULL PRIMARY KEY,
                          hashtag text
);

CREATE TABLE jobs (
                      id bigserial NOT NULL PRIMARY KEY,
                      job text
);

CREATE TABLE matches (
                         id bigserial NOT NULL PRIMARY KEY,
                         user_first_id bigint references users(id),
                         user_second_id bigint references users(id),
                         shown boolean
);


CREATE TABLE reactions (
                           id bigserial NOT NULL PRIMARY KEY,
                           reaction text
);


CREATE TABLE reasons (
                         id bigserial NOT NULL PRIMARY KEY,
                         reason text
);

CREATE TABLE sexes (
                       id bigserial NOT NULL PRIMARY KEY,
                       sex text
);


CREATE TABLE statuses (
                          id bigserial NOT NULL PRIMARY KEY,
                          status text,
                          max_like bigint,
                          advertising boolean
);

CREATE TABLE zodiacs (
                         id bigserial NOT NULL PRIMARY KEY,
                         zodiac text
);


-- user_infos, user_filters, user_photos -
-- эти таблицы соответствуют 3NF.
-- Все атрибуты функционально зависят от первичного ключа и нет транзитивных зависимостей.

CREATE TABLE user_filters (
                              id bigserial NOT NULL PRIMARY KEY,
                              user_id bigint UNIQUE references users(id),
                              min_age bigint,
                              max_age bigint,
                              search_sex bigint references sexes(id)
);

CREATE TABLE user_infos (
                            id bigserial NOT NULL PRIMARY KEY,
                            user_id bigint UNIQUE references users(id),
                            name text,
                            city_id bigint,
                            sex bigint references sexes(id),
                            description text,
                            zodiac bigint references zodiacs(id),
                            job bigint references jobs(id),
                            education bigint references educations(id)
);

CREATE TABLE user_photos (
                             id bigserial NOT NULL PRIMARY KEY,
                             user_id bigint references users(id),
                             photo bigint references media(id),
                             avatar boolean
);

-- user_hashtags, user_histories, user_reactions, user_reasons, chat_users -
-- эти таблицы соответствуют 3NF.
-- Они представляют собой связующие таблицы между другими таблицами.
-- Все атрибуты функционально зависят от первичного ключа и нет транзитивных зависимостей.

CREATE TABLE user_hashtags (
                               id bigserial NOT NULL PRIMARY KEY,
                               user_id bigint references users(id),
                               hashtag_id bigint references hashtags(id)
);

CREATE TABLE user_histories (
                                id bigserial NOT NULL PRIMARY KEY,
                                user_id bigint references users(id) NOT NULL,
                                user_profile_id bigint references users(id) NOT NULL ,
                                show_date date
);


CREATE TABLE user_reactions (
                                id bigserial NOT NULL PRIMARY KEY,
                                user_id bigint references users(id),
                                user_from_id bigint references users(id),
                                reaction_id bigint references reactions(id)
);

CREATE TABLE user_reasons (
                              id bigserial NOT NULL PRIMARY KEY,
                              user_id bigint references users(id),
                              reason_id bigint references reasons(id)
);

CREATE TABLE chat_users (
                            id bigserial NOT NULL PRIMARY KEY,
                            chat_id bigint references chats(id) NOT NULL,
                            user_id bigint references users(id) NOT NULL
);

-- users - эта таблица соответствует 3NF.
-- Все атрибуты функционально зависят от id, которое является первичным ключом.
-- Нет транзитивных зависимостей.
CREATE TABLE users (
                       id bigserial NOT NULL PRIMARY KEY,
                       email text UNIQUE,
                       password_hash text,
                       birth_day date,
                       status bigint references statuses(id) DEFAULT 1 NOT NULL,
                       count bigint DEFAULT 0 NOT NULL
);


-- chats, files, messages, media - эти таблицы соответствуют 3NF.
-- Все атрибуты функционально зависят от первичного ключа и нет транзитивных зависимостей.

CREATE TABLE chats (
    id bigserial NOT NULL PRIMARY KEY
);

CREATE TABLE files (
                       id bigserial NOT NULL PRIMARY KEY,
                       path text NOT NULL,
                       user_id bigint  references users(id) NOT NULL
);

CREATE TABLE media (
                       id bigserial NOT NULL PRIMARY KEY,
                       message_id bigint references messages(id) NOT NULL,
                       message_type text NOT NULL,
                       path text references files(path) NOT NULL
);

CREATE TABLE messages (
                          id bigserial NOT NULL PRIMARY KEY,
                          chat_id bigint references chats(id) NOT NULL,
                          sender_id bigint references users(id) NOT NULL,
                          content text NOT NULL,
                          sent_at timestamp without time zone NOT NULL,
                          read_status boolean NOT NULL DEFAULT false
);
