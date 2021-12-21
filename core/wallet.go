package core

import (
	"context"
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

var (
	ErrInvalidTrace          = errors.New("invalid trace")
	ErrInvalidConversationID = errors.New("invalid conversation id")
)

type (
	Asset struct {
		AssetID   string          `sql:"type:char(36)" gorm:"primaryKey" json:"asset_id"`
		Name      string          `sql:"type:varchar(64)" json:"name"`
		ChainID   string          `sql:"type:char(36)" json:"chain_id"`
		PriceUSD  decimal.Decimal `sql:"type:decimal(64,8)" json:"price_usd"`
		PriceBTC  decimal.Decimal `sql:"type:decimal(64,8)" json:"price_btc"`
		Symbol    string          `sql:"type:varchar(32)" gorm:"index:idx_symbol" json:"symbol"`
		IconURL   string          `sql:"type:varchar(1024)" json:"icon_url"`
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
	}

	Snapshot struct {
		ID              uint64          `gorm:"primaryKey" json:"id,omitempty"`
		SnapshotID      string          `sql:"type:char(36)" gorm:"index:idx_snapshots_snapshot_id;unique" json:"snapshot_id,omitempty"`
		TraceID         string          `sql:"type:char(36)" json:"trace_id,omitempty"`
		Source          string          `sql:"type:varchar(32)" json:"source,omitempty"`
		TransactionHash string          `sql:"type:varchar(64)" json:"transaction_hash,omitempty"`
		Receiver        string          `sql:"type:varchar(256)" json:"receiver,omitempty"`
		Sender          string          `sql:"type:varchar(256)" json:"sender,omitempty"`
		Type            string          `sql:"type:varchar(32)" json:"type,omitempty"`
		CreatedAt       time.Time       `json:"created_at,omitempty"`
		UserID          string          `sql:"type:char(36)" json:"user_id,omitempty"`
		OpponentID      string          `sql:"type:char(36)" json:"opponent_id,omitempty"`
		AssetID         string          `sql:"type:char(36)" json:"asset_id,omitempty"`
		Amount          decimal.Decimal `sql:"type:decimal(64,8)" json:"amount,omitempty"`
		Memo            string          `sql:"type:varchar(256)" gorm:"default:''" json:"memo,omitempty"`
	}

	WalletStore interface {
		GetSnapshots(ctx context.Context, userID string, from time.Time, limit int, assetID string) ([]*Snapshot, error)
		GetSnapshot(ctx context.Context, snapshotID string) (Snapshot, error)
		SetSnapshots(ctx context.Context, snapshots []*Snapshot) error
		GetAssets(ctx context.Context) ([]*Asset, error)
		GetAsset(ctx context.Context, assetID string) (*Asset, error)
		SetAssets(ctx context.Context, assets []*Asset) error
		PollSnapshots(ctx context.Context, offset time.Time, limit int) ([]*Snapshot, error)
	}
)
