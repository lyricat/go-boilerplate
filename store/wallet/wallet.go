package wallet

import (
	"context"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pkg/store/db"
	core "github.com/lyricat/go-boilerplate/core/wallet"
)

func init() {

	db.RegisterMigrate(func(db *db.DB) error {
		tx := db.Update().Model(core.Asset{})
		if err := tx.AutoMigrate(core.Asset{}).Error; err != nil {
			return err
		}
		return nil
	})

	db.RegisterMigrate(func(db *db.DB) error {
		tx := db.Update().Model(core.Snapshot{})
		if err := tx.AutoMigrate(core.Snapshot{}).Error; err != nil {
			return err
		}
		return nil
	})

}

func New(db *db.DB, client *mixin.Client, cfg Config) *WalletStore {
	return &WalletStore{db: db, client: client, cfg: cfg}
}

type Config struct {
	Pin string `valid:"required"`
}

type WalletStore struct {
	db     *db.DB
	client *mixin.Client
	cfg    Config
}

// type Store struct {
// 	db     *db.DB
// 	client *mixin.Client
// 	cfg    core.Config
// }

func (s *WalletStore) GetSnapshots(ctx context.Context, userID string, from time.Time, limit int, assetID string) ([]*core.Snapshot, error) {
	var ss []*core.Snapshot
	query := s.db.View().Limit(limit)

	if assetID == "" {
		query = query.Where("opponent_id = ?", userID)
	} else {
		query = query.Where("opponent_id = ? AND asset_id = ?", userID, assetID)
	}
	if !from.IsZero() {
		query = query.Where("created_at < ?", from)
	}
	err := query.Order("created_at DESC").Find(&ss).Error
	return ss, err
}

func (s *WalletStore) GetSnapshot(ctx context.Context, snapshotID string) (core.Snapshot, error) {
	var ss core.Snapshot
	err := s.db.View().Where("snapshot_id = ?", snapshotID).Find(&ss).Error
	return ss, err
}

func (s *WalletStore) SetSnapshots(ctx context.Context, snapshots []*core.Snapshot) error {
	return s.db.Tx(func(tx *db.DB) error {
		for _, snapshot := range snapshots {
			if err := tx.Update().Where("snapshot_id = ?", snapshot.SnapshotID).FirstOrCreate(snapshot).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *WalletStore) GetAssets(ctx context.Context) ([]*core.Asset, error) {
	var assets []*core.Asset
	err := s.db.View().Find(&assets).Error
	return assets, err
}

func (s *WalletStore) GetAsset(ctx context.Context, assetID string) (*core.Asset, error) {
	var asset core.Asset
	err := s.db.View().Where("asset_id = ?", assetID).Find(&asset).Error
	return &asset, err
}

func (s *WalletStore) SetAssets(ctx context.Context, assets []*core.Asset) error {
	return s.db.Tx(func(tx *db.DB) error {
		for _, asset := range assets {
			if err := tx.Update().Where("asset_id = ?", asset.AssetID).FirstOrCreate(asset).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
