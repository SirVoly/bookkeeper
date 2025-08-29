# Helpfull commands

## Run project
go run .

## Goose
goose -dir sql/schema sqlite3 bookkeeper.db up
goose -dir sql/schema sqlite3 bookkeeper.db down

## SQLC
sqlc generate

## Cobra-CLI
cobra-cli add <commandName>