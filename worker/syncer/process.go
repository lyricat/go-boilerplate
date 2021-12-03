package syncer

import (
	"context"

	"github.com/fox-one/pkg/logger"
	"github.com/lyricat/go-boilerplate/core/wallet"
)

func (w *Worker) ProcessSnapshots(ctx context.Context, snapshots []*wallet.Snapshot) error {
	for _, snapshot := range snapshots {
		err := w.ProcessSnapshot(ctx, snapshot)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) ProcessSnapshot(ctx context.Context, snapshot *wallet.Snapshot) error {
	log := logger.FromContext(ctx)
	if err := w.PersistentSnapshot(ctx, snapshot); err != nil {
		log.Warn("failed to persistent snapshot", err)
		return err
	}

	return nil
}

func (w *Worker) PersistentSnapshot(ctx context.Context, snapshot *wallet.Snapshot) (err error) {
	if err := w.wallets.SetSnapshots(ctx, []*wallet.Snapshot{snapshot}); err != nil {
		return err
	}
	return nil
}
