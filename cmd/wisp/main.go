package main

import (
	"fmt"

	"github.com/joatisio/wisp/internal/app"
)

func main() {
	c := setupConfig()

	dsn := generateDsn(c.Database)
	db := setupDatabase(createDialector(dsn, nil))

	logger := setupLogger(c.Logger)

	ap := app.NewApp(c, db, logger)
	fmt.Println(ap)
}
