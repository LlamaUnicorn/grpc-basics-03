-- +goose Up
create table link (
    id serial primary key,
    url text not null,
    title text,
    description text,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table link;