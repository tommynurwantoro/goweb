package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/tommynurwantoro/goweb/internal/adapter/rest"
	"github.com/tommynurwantoro/goweb/views/pages"
)

type LoginHandler struct {
	App *rest.Fiber `inject:"fiber"`
}

func (h *LoginHandler) Startup() error {
	h.App.Get("/login", h.Login)

	return nil
}

func (h *LoginHandler) Shutdown() error { return nil }

func (h *LoginHandler) Login(c *fiber.Ctx) error {
	handler := adaptor.HTTPHandler(templ.Handler(pages.Login()))
	return handler(c)
}
