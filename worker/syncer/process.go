package syncer

import (
	"context"

	"github.com/fox-one/pkg/logger"
	"github.com/lyricat/go-boilerplate/core"
)

func (w *Worker) ProcessSnapshots(ctx context.Context, snapshots []*core.Snapshot) error {
	for _, snapshot := range snapshots {
		err := w.ProcessSnapshot(ctx, snapshot)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) ProcessSnapshot(ctx context.Context, snapshot *core.Snapshot) error {
	log := logger.FromContext(ctx)
	if err := w.PersistentSnapshot(ctx, snapshot); err != nil {
		log.Warn("failed to persistent snapshot", err)
		return err
	}

	return nil
}

func (w *Worker) PersistentSnapshot(ctx context.Context, snapshot *core.Snapshot) (err error) {
	if err := w.wallets.SetSnapshots(ctx, []*core.Snapshot{snapshot}); err != nil {
		return err
	}
	return nil
}
