package bootstrap

import "github.com/tommynurwantoro/goweb/internal/adapter/rest"

func (b *Bootstrap) RegisterRest() {
	b.Container.RegisterService("fiber", new(rest.Fiber))
}
