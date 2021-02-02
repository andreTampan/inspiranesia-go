package main

import (
	"context"
	"inspiranesia/system/config"
	"inspiranesia/system/http"
	"inspiranesia/system/logging"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(logging.ProvideLogging),
		fx.Provide(config.ProvideViper),
		fx.Provide(http.ProvideHttpHandler),
		fx.Invoke(start),
	)

	app.Run()
}

func start(lifecycle fx.Lifecycle, logger logging.NougatLoggingProvider, handler http.NougatHttpHandler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go handler.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
}
