package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/tommynurwantoro/goweb/internal/adapter/rest"
	"github.com/tommynurwantoro/goweb/views/pages"
)

type HomeHandler struct {
	App *rest.Fiber `inject:"fiber"`
}

func (h *HomeHandler) Startup() error {
	h.App.Get("/", h.Home)

	return nil
}

func (h *HomeHandler) Shutdown() error { return nil }

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	handler := adaptor.HTTPHandler(templ.Handler(pages.Home()))
	return handler(c)
}
