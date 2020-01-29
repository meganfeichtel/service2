package handlers

import (
	"log"
	"os"

	"github.com/meganfeichtel/service2/internal/platform/web"
)

//API comment
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown)

	app.Handle("GET", "/test", health)
	return app
}
