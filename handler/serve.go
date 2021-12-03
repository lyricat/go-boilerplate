package handler

import (
	"fmt"
	"net/http"

	"github.com/lyricat/go-boilerplate/handler/asset"
	"github.com/lyricat/go-boilerplate/handler/echo"
	"github.com/lyricat/go-boilerplate/handler/render"
	"github.com/lyricat/go-boilerplate/store/wallet"

	"github.com/go-chi/chi"
)

func New(cfg Config, wallets *wallet.WalletStore) Server {
	return Server{cfg: cfg, wallets: wallets}
}

type (
	Config struct {
	}

	Server struct {
		cfg     Config
		wallets *wallet.WalletStore
	}
)

func (s Server) HandleRest() http.Handler {
	r := chi.NewRouter()
	r.Use(render.WrapResponse(true))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, http.StatusNotFound, fmt.Errorf("not found"))
	})

	r.Route("/echo", func(r chi.Router) {
		r.Get("/{msg}", echo.Get())
		r.Post("/", echo.Post())
	})

	r.Route("/assets", func(r chi.Router) {
		r.Get("/", asset.GetAssets(s.wallets))
		r.Get("/{assetID}", asset.GetAsset(s.wallets))
	})

	return r
}
