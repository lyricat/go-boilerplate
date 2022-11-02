package asset

import (
	"database/sql"
	"net/http"

	"go-boilerplate/core"
	"go-boilerplate/handler/render"

	"github.com/go-chi/chi"
)

func GetAsset(store core.AssetStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assetID := chi.URLParam(r, "assetID")
		asset, err := store.GetAsset(ctx, assetID)
		if err == sql.ErrNoRows {
			render.Error(w, http.StatusNotFound, err)
			return
		}
		if err != nil {
			render.Error(w, http.StatusInternalServerError, err)
			return
		}
		render.JSON(w, asset)
	}
}

func GetAssets(store core.AssetStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assets, err := store.GetAssets(ctx)
		if err != nil {
			render.Error(w, http.StatusInternalServerError, err)
		}
		render.JSON(w, assets)
	}
}
