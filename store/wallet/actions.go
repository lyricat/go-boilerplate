package wallet

import (
	"context"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	core "github.com/lyricat/go-boilerplate/core"
)

func (s *WalletStore) PollSnapshots(ctx context.Context, offset time.Time, limit int) ([]*core.Snapshot, error) {
	snapshots, err := s.client.ReadNetworkSnapshots(ctx, "", offset, "ASC", limit)
	if err != nil {
		return nil, err
	}

	return convertSnapshots(snapshots), nil
}

func convertSnapshots(items []*mixin.Snapshot) []*core.Snapshot {
	var snapshots = make([]*core.Snapshot, len(items))
	for i, s := range items {
		snapshots[i] = &core.Snapshot{
			CreatedAt:       s.CreatedAt,
			SnapshotID:      s.SnapshotID,
			UserID:          s.UserID,
			OpponentID:      s.OpponentID,
			TraceID:         s.TraceID,
			AssetID:         s.AssetID,
			Amount:          s.Amount,
			Memo:            s.Memo,
			Receiver:        s.Receiver,
			Sender:          s.Sender,
			Type:            s.Type,
			Source:          s.Source,
			TransactionHash: s.TransactionHash,
		}
	}
	return snapshots
}
