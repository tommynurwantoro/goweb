package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/runsystemid/gontainer"
	"github.com/tommynurwantoro/goweb/config"
	"github.com/tommynurwantoro/goweb/internal/adapter/rest"
)

type Bootstrap struct {
	Container gontainer.Container
}

func New() *Bootstrap {
	return &Bootstrap{
		Container: gontainer.New(),
	}
}

func (b *Bootstrap) Run(conf *config.Model) {
	b.Container.RegisterService("config", conf)

	b.RegisterRest()
	b.RegisterAPI()

	// Startup the container
	if err := b.Container.Ready(); err != nil {
		log.Panic("Failed to populate service", err)
	}

	// Start server
	fiberApp := b.Container.GetServiceOrNil("fiber").(*rest.Fiber)
	errs := make(chan error, 2)
	go func() {
		log.Info("Listening on port : ", conf.Http.Port)
		errs <- fiberApp.Listen(fmt.Sprintf(":%d", conf.Http.Port))
	}()

	log.Info("Server started")
	b.gracefulShutdown()
}

func (b *Bootstrap) gracefulShutdown() {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	delay := 1 * time.Second

	log.Info(fmt.Sprintf("Signal termination received. Waiting %v to shutdown.", delay))
	time.Sleep(delay)

	log.Info("Cleaning up resources...")

	// This will shuting down all the resources
	b.Container.Shutdown()

	log.Info("Bye")
}
