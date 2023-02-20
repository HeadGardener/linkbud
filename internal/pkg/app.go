package app

import (
	"github.com/HeadGardener/linkbud/configs"
	postgres "github.com/HeadGardener/linkbud/internal/app/db"
	"github.com/HeadGardener/linkbud/internal/app/handlers"
	"github.com/HeadGardener/linkbud/internal/app/repository"
	"github.com/HeadGardener/linkbud/internal/app/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	g *gin.Engine
}

func New(conf configs.DBConfig) (*App, error) {
	app := &App{}

	db, err := postgres.NewDB(conf)
	if err != nil {
		return nil, err
	}
	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	app.g = handler.InitRoutes()

	return app, nil
}

func (app *App) Start(conf configs.ServerConfig) error {
	return app.g.Run(":" + conf.Port)
}
