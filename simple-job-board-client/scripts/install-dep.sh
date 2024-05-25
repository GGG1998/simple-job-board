#!/bin/bash

cd internal/tools
go get github.com/golang-migrate/migrate/v4/cmd/migrate
go install github.com/golang-migrate/migrate/v4/cmd/migrate
go get github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
