package tools

import (
	"time"

	"github.com/hhelix/baulkham/internal/config"
)

// ping database
func PingDB(app *config.Application) error {
	start := time.Now()

	db, err := app.DB.DB()
	if err != nil {
		return err
	}
	app.Logger.Info("Pinging database...")
	err = db.Ping()
	if err != nil {
		return err
	}

	elapsed := time.Since(start)
	app.Logger.Infof("Database pinged successfully in %s", elapsed)

	return nil
}
