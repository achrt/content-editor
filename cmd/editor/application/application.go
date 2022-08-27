package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"

	"content-editor/internal/domain/repositories"
	"content-editor/internal/domain/repositories/catalogue"
	"content-editor/internal/infrastracture/db"
)

type App struct {
	Router    *fiber.App
	Log       log.Logger
	Catalogue repositories.Catalogue

	address string
}

func New(router *fiber.App) (*App, error) {
	cfg := Config{}
	if err := cfg.loadConfiguration(); err != nil {
		return nil, err
	}

	dbc, sqlDb, err := db.DBConnect(&cfg.DBConfigSection)
	if err != nil {
		return nil, err
	}
	defer sqlDb.Close()

	app := &App{
		Router: router,
	}

	app.Catalogue = catalogue.New(dbc)
	return app, nil
}

func (a *App) Run() error {
	a.Log.Infof("starting service at: http://%s", a.address)
	return a.Router.Listen(a.address)
}
