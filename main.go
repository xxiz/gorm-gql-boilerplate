package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hhelix/baulkham/internal/config"
	"github.com/hhelix/baulkham/internal/handlers"
	"github.com/hhelix/baulkham/internal/routes"
	"github.com/hhelix/baulkham/internal/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// Logging Method Name
	//log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	flag.IntVar(&config.App.Config.Port, "port", 5000, "Server port to listen on")
	flag.StringVar(&config.App.Config.Env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&config.App.Config.Db.Dsn, "db-dsn", "postgres://baulkham:baulkham@localhost:5433/postgres?sslmode=disable", "Postgres Database connection string")
	flag.StringVar(&config.App.Config.Jwt.Secret, "jwt-secret", "", "secret")
	config.App.Logger = log.StandardLogger()
}

func main() {
	flag.Parse()
	config.App.Logger.Info("Starting server on port ", config.App.Config.Port)
	db, err := gorm.Open(postgres.Open(config.App.Config.Db.Dsn), &gorm.Config{})
	if err != nil {
		config.App.Logger.Error(err)
		panic("Failed to connect to database")
	}
	config.App.DB = db
	// Ping database
	tools.PingDB(&config.App)

	repo := handlers.NewRepo(&config.App, db)
	handlers.NewHandlers(repo)

	config.App.Gin = routes.Routes(&config.App)
	err = config.App.Gin.Run(fmt.Sprintf(":%d", config.App.Config.Port))
	if err != nil {
		config.App.Logger.Fatal(err)
	}
}
