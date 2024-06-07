# /pets
Основное приложение, рест апи, валидация подключения, миграции, докер


# /pets-clicker
Websocket сервер для мини игры кликер


# Migrations 

## Goose
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest 
```
```shell
export GOOSE_DRIVER=postgres
export GOOSE_MIGRATION_DIR=./migrations
export GOOSE_DBSTRING=postgres://postgres:1234@localhost:5444/pets
```
