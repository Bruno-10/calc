// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	v1 "github.com/Bruno-10/calc/app/services/api/handlers/v1"
	"github.com/Bruno-10/calc/business/web/v1/mid"
	"github.com/Bruno-10/calc/foundation/logger"
	"github.com/Bruno-10/calc/foundation/web"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	if opts.corsOrigin != "" {
		app.EnableCORS(mid.Cors(opts.corsOrigin))
	}

	v1.Routes(app, v1.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
	})

	return app
}
