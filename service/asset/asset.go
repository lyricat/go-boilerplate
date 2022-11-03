package asset

import (
	"context"
	"go-boilerplate/core"

	"github.com/fox-one/mixin-sdk-go"
)

func New(client *mixin.Client, assets core.AssetStore) *assetService {
	return &assetService{client: client, assets: assets}
}

type assetService struct {
	client *mixin.Client
	assets core.AssetStore
}

func (s *assetService) UpdateAssets(ctx context.Context) error {
	as, err := s.assets.GetAssets(ctx)
	if err != nil {
		return err
	}
	var newAs []*mixin.Asset
	for _, a := range as {
		ma, err := s.client.ReadAsset(ctx, a.AssetID)
		if err != nil {
			return err
		}
		newAs = append(newAs, ma)
	}

	return s.assets.SetAssets(ctx, convertAssets(newAs))
}

func convertAssets(items []*mixin.Asset) []*core.Asset {
	var assets = make([]*core.Asset, len(items))
	for i, item := range items {
		assets[i] = convertAsset(item)
	}
	return assets
}

func convertAsset(item *mixin.Asset) *core.Asset {
	return &core.Asset{
		AssetID:  item.AssetID,
		ChainID:  item.ChainID,
		Name:     item.Name,
		Symbol:   item.Symbol,
		IconURL:  item.IconURL,
		PriceUSD: item.PriceUSD,
	}
}
