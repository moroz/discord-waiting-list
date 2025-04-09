-- +goose Up
-- +goose StatementBegin
create table waiting_list (
  id integer primary key autoincrement,
  email text not null unique,
  name text not null,
  bio text,
  region text,
  inserted_at bigint not null default (unixepoch()),
  updated_at bigint not null default (unixepoch()),
  check (region is null or length(region) = 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table waiting_list;
-- +goose StatementEnd
