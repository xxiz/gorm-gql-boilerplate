package dbrepo

import (
	"github.com/hhelix/baulkham/internal/config"
	"gorm.io/gorm"
)

type postgresDBRepo struct {
	App *config.Application
	DB  *gorm.DB
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(a *config.Application, db *gorm.DB) *postgresDBRepo {
	return &postgresDBRepo{
		App: a,
		DB:  db,
	}
}
