package asset

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/lyricat/go-boilerplate/handler/render"
	"github.com/lyricat/go-boilerplate/store/wallet"
)

func GetAsset(store *wallet.WalletStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assetID := chi.URLParam(r, "assetID")
		asset, err := store.GetAsset(ctx, assetID)
		if err == gorm.ErrRecordNotFound {
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

func GetAssets(store *wallet.WalletStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assets, err := store.GetAssets(ctx)
		if err != nil {
			render.Error(w, http.StatusInternalServerError, err)
		}
		render.JSON(w, assets)
	}
}
