package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)

//FOUNDATIONAL CODE - not allowed to log

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	//app is now a mux
	*httptreemux.TreeMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	app := App{
		TreeMux:  httptreemux.New(),
		shutdown: shutdown,
	}
	return &app
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		if err := handler(r.Context(), w, r, params); err != nil {
			return
		}
	}
	a.TreeMux.Handle(verb, path, h)
}
