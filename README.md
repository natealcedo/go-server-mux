# README

This is an exploration of creating an api server using go 1.22 
and it's improvements stated here

https://go.dev/blog/routing-enhancements


## Migrations
This repository uses [goose](https://github.com/pressly/goose) to handle migrations

You can install using go install, homebrew or whatever wish you like

To run migrations, you can use the following command
```bash
make up
```

To create a migration run the following command
```bash
make create name=name_of_migration
```
