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

	DefaultSnapshotFetchCount = 100
)

type (
	Snapshot struct {
		SnapshotID      string          `json:"snapshot_id,omitempty"`
		TraceID         string          `json:"trace_id,omitempty"`
		Source          string          `json:"source,omitempty"`
		TransactionHash string          `json:"transaction_hash,omitempty"`
		Receiver        string          `json:"receiver,omitempty"`
		Sender          string          `json:"sender,omitempty"`
		Type            string          `json:"type,omitempty"`
		CreatedAt       time.Time       `json:"created_at,omitempty"`
		UserID          string          `json:"user_id,omitempty"`
		OpponentID      string          `json:"opponent_id,omitempty"`
		AssetID         string          `json:"asset_id,omitempty"`
		Amount          decimal.Decimal `json:"amount,omitempty"`
		Memo            string          `json:"memo,omitempty"`
	}

	SnapshotStore interface {
		GetSnapshots(ctx context.Context, from time.Time, limit int) ([]*Snapshot, error)
		GetSnapshot(ctx context.Context, snapshotID string) (*Snapshot, error)
		SetSnapshots(ctx context.Context, snapshots []*Snapshot) error
	}
	SnapshotService interface {
		PollSnapshots(ctx context.Context, offset time.Time, limit int) ([]*Snapshot, error)
	}
)
