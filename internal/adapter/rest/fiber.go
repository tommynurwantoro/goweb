package rest

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"github.com/tommynurwantoro/goweb/config"
)

type Fiber struct {
	*fiber.App
	Conf *config.Model `inject:"config"`
}

func (f *Fiber) Startup() error {
	//starting http
	f.App = fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(f.Conf.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(f.Conf.Http.WriteTimeout) * time.Second,
		Views:        html.NewFileSystem(http.Dir("./public"), ".html"),
		ViewsLayout:  "main",
	})

	f.App.Static("/static", "./static")

	f.App.Use(requestid.New())
	f.App.Use(logger.New())

	return nil
}

func (f *Fiber) Shutdown() error { return f.App.Shutdown() }
