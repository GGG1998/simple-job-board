//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/sqlc-dev/sqlc"
)
