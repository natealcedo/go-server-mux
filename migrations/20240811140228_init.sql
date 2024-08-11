-- +goose Up
-- +goose StatementBegin
create table test_table (
  id uuid primary key default gen_random_uuid(),
  name text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table test_table;
-- +goose StatementEnd
