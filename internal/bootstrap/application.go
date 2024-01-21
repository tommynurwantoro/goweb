package bootstrap

import "github.com/tommynurwantoro/goweb/internal/application/api"

func (b *Bootstrap) RegisterAPI() {
	// Registering API
	b.Container.RegisterService("home", new(api.HomeHandler))
	b.Container.RegisterService("login", new(api.LoginHandler))
}
