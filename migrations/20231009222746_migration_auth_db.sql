-- +goose Up
create table chats (
    id serial primary key,
    username text not null,
    msg text not null
);

-- +goose Down
drop table chats;
