-- +goose Up
create table chat_server (
    id serial primary key,
    msg text not null
);

-- +goose Down
drop table chat_server;
