package wallet

import (
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
		AssetID string          `gorm:"primaryKey" json:"asset_id"`
		Name    string          `json:"name"`
		Symbol  string          `gorm:"index:idx_symbol" json:"symbol"`
		Logo    string          `json:"icon_url"`
		Balance decimal.Decimal `sql:"type:decimal(64,8)" json:"balance"`
	}

	Snapshot struct {
		ID              uint64          `gorm:"primaryKey" json:"id,omitempty"`
		SnapshotID      string          `gorm:"required,index:idx_snapshots_snapshot_id;unique" json:"snapshot_id,omitempty"`
		TraceID         string          `gorm:"required" json:"trace_id,omitempty"`
		Source          string          `gorm:"required" json:"source,omitempty"`
		TransactionHash string          `gorm:"required" json:"transaction_hash,omitempty"`
		Receiver        string          `gorm:"required" json:"receiver,omitempty"`
		Sender          string          `gorm:"required" json:"sender,omitempty"`
		Type            string          `gorm:"required" json:"type,omitempty"`
		CreatedAt       time.Time       `gorm:"required" json:"created_at,omitempty"`
		UserID          string          `gorm:"required" json:"user_id,omitempty"`
		OpponentID      string          `gorm:"required" json:"opponent_id,omitempty"`
		AssetID         string          `gorm:"required" json:"asset_id,omitempty"`
		Amount          decimal.Decimal `sql:"type:decimal(64,8)" gorm:"required" json:"amount,omitempty"`
		Memo            string          `gorm:"required,default:''" json:"memo,omitempty"`
	}
)
