package core

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type (
	Asset struct {
		AssetID   string          `json:"asset_id"`
		Name      string          `json:"name"`
		ChainID   string          `json:"chain_id"`
		PriceUSD  decimal.Decimal `json:"price_usd"`
		Symbol    string          `json:"symbol"`
		IconURL   string          `json:"icon_url"`
		UpdatedAt time.Time       `json:"updated_at"`
	}

	AssetStore interface {
		GetAssets(ctx context.Context) ([]*Asset, error)
		GetAsset(ctx context.Context, assetID string) (*Asset, error)
		SetAssets(ctx context.Context, assets []*Asset) error
	}
)
