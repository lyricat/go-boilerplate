package handler

import (
	"net/http"

	"github.com/lyricat/go-boilerplate/handler/echo"
	"github.com/lyricat/go-boilerplate/handler/render"

	"github.com/go-chi/chi"
	"github.com/twitchtv/twirp"
)

func New(cfg Config) Server {
	return Server{cfg: cfg}
}

type (
	Config struct {
	}

	Server struct {
		cfg Config
	}
)

func (s Server) HandleRest() http.Handler {
	r := chi.NewRouter()
	r.Use(render.WrapResponse(true))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, twirp.NotFoundError("not found"))
	})

	r.Route("/echo", func(r chi.Router) {
		r.Get("/{msg}", echo.HandleGet())
		r.Post("/", echo.HandlePost())
	})

	return r
}

func resetRoutePath(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if c := chi.RouteContext(ctx); c != nil {
			c.RoutePath = r.URL.Path
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
