package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hhelix/baulkham/internal/config"
	"github.com/hhelix/baulkham/internal/graphql"
	"github.com/hhelix/baulkham/internal/repository"
	"github.com/hhelix/baulkham/internal/repository/dbrepo"
	"gorm.io/gorm"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DatabaseRepo
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// NewRepo creates a new Repository
func NewRepo(a *config.Application, db *gorm.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(a, db),
	}
}

// GraphqlHandler Defining the Graphql handler
func (m *Repository) GraphqlHandler() gin.HandlerFunc {
	var mb int64 = 1 << 20
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := graphql.Config{Resolvers: &graphql.Resolver{App: m.App, DB: m.DB}}

	//TODO: Add directives
	// c.Directives.Auth = directives.Auth

	h := handler.NewDefaultServer(graphql.NewExecutableSchema(c))
	// For File upload
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{
		MaxUploadSize: 32 * mb,
		MaxMemory:     50 * mb,
	})
	h.Use(extension.Introspection{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (m *Repository) PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
