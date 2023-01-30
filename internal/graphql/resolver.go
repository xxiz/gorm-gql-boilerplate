package graphql

import (
	"github.com/hhelix/baulkham/internal/config"
	"github.com/hhelix/baulkham/internal/repository"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App *config.Application
	DB  repository.DatabaseRepo
}
