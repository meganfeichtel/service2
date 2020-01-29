package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/meganfeichtel/service2/internal/platform/web"
)

//Logger comment
func Logger(log *log.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			handler(ctx, w, r, params)
			return nil
		}

		//LOGGGGGGG
		log.Println("lololololololol")

		return h
	}

	return m
}
