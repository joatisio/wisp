package main

import (
	"fmt"

	"github.com/joatisio/wisp/internal/app"
)

const APIV1Prefix = "/packages/v1/"

func main() {
	c := setupConfig()

	dsn := generateDsn(c.Database)
	db := setupDatabase(createDialector(dsn, nil))

	logger := setupLogger(c.Logger)

	engine := setupWebServer(c.Server, logger)

	srv := app.NewServer(engine, logger, c.Server, APIV1Prefix)

	appInstance := app.NewApp(c, db, logger, srv)

	setupRoutes(appInstance)

	srv.Serve(c.Server.AddrString(), c.Server.)
}
