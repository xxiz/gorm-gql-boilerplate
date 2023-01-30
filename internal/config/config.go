package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const Version = "0.0.1"

var App Application

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Jwt struct {
		Secret string
	}
}

type Application struct {
	Config Config
	Logger *logrus.Logger
	DB     *gorm.DB
	Gin    *gin.Engine
}
