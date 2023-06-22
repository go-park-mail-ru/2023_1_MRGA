CREATE TABLE files (
    id bigserial NOT NULL PRIMARY KEY,
    path text NOT NULL,
    user_id bigint  references users(id) NOT NULL
);