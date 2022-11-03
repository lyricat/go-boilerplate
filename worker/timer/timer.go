package timer

import (
	"context"
	"time"

	"go-boilerplate/core"

	"github.com/asaskevich/govalidator"
	"github.com/fox-one/pkg/logger"
)

type (
	Config struct {
	}

	Worker struct {
		cfg       Config
		propertys core.PropertyStore
		assetz    core.AssetService
	}
)

const (
	updateAssetsCheckpoint = "timer:update_assets_checkpoint"
)

func New(
	cfg Config,
	propertys core.PropertyStore,
	assetz core.AssetService,
) *Worker {
	if _, err := govalidator.ValidateStruct(cfg); err != nil {
		panic(err)
	}

	return &Worker{
		cfg:       cfg,
		propertys: propertys,
		assetz:    assetz,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	log := logger.FromContext(ctx).WithField("worker", "timer")
	ctx = logger.WithContext(ctx, log)

	dur := time.Millisecond
	var circle int64
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(dur):
			if err := w.run(ctx, circle); err == nil {
				dur = time.Second
				circle += 1
			} else {
				dur = 10 * time.Second
				circle += 10
			}
		}
	}
}

func (w *Worker) run(ctx context.Context, circle int64) error {
	if err := w.updateAssets(ctx); err != nil {
		return err
	}

	return nil
}

func (w *Worker) updateAssets(ctx context.Context) error {
	log := logger.FromContext(ctx).WithField("worker", "timer")

	v, err := w.propertys.Get(ctx, updateAssetsCheckpoint)
	if err != nil {
		log.Warn("failed to get property", "err", err)
		return err
	}

	var (
		lastInvokedAt = v.Time()
		now           = time.Now()
	)

	if now.Sub(lastInvokedAt) > 3*time.Minute {
		log.Println("start invoke", now)
		// update assets and fiats
		if err := w.assetz.UpdateAssets(ctx); err != nil {
			log.Warn("update assets error", err)
		}

		if err := w.propertys.Set(ctx, updateAssetsCheckpoint, now.Format(time.RFC3339Nano)); err != nil {
			log.WithError(err).Errorln("propertys.Save")
			return err
		}
		log.Println("end of invoking", time.Now())
	}
	return nil
}
