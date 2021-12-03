package worker

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fox-one/pkg/logger"
	"github.com/fox-one/pkg/store/db"
	propertystore "github.com/fox-one/pkg/store/property"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lyricat/go-boilerplate/config"
	"github.com/lyricat/go-boilerplate/handler/hc"
	"github.com/lyricat/go-boilerplate/session"
	"github.com/lyricat/go-boilerplate/store/wallet"
	"github.com/lyricat/go-boilerplate/worker"
	"github.com/lyricat/go-boilerplate/worker/messenger"
	"github.com/lyricat/go-boilerplate/worker/syncer"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"

	"github.com/spf13/cobra"
)

func NewCmdWorker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker [health check port]",
		Short: "run workers",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ctx := cmd.Context()
			s := session.From(ctx)

			database := db.MustOpen(config.C().DB)
			defer database.Close()

			property := propertystore.New(database)

			keystore, err := s.GetKeystore()
			if err != nil {
				return err
			}

			client, err := s.GetClient()
			if err != nil {
				return err
			}

			pin, err := s.GetPin()
			if err != nil {
				return err
			}

			wallets := wallet.New(database, client, wallet.Config{Pin: pin})

			workers := []worker.Worker{
				// syncer
				syncer.New(syncer.Config{
					ClientID: keystore.ClientID,
				}, property, wallets),
				// messenger
				messenger.New(client),
			}

			// run them all
			g, ctx := errgroup.WithContext(ctx)
			for idx := range workers {
				w := workers[idx]
				g.Go(func() error {
					return w.Run(ctx)
				})
			}

			// start the health check server
			g.Go(func() error {
				mux := chi.NewMux()
				mux.Use(middleware.Recoverer)
				mux.Use(middleware.StripSlashes)
				mux.Use(cors.AllowAll().Handler)
				mux.Use(logger.WithRequestID)
				mux.Use(middleware.Logger)

				{
					// hc for workers
					mux.Mount("/hc", hc.Handle(cmd.Version))
				}

				// launch server
				port := 8081
				if len(args) > 0 {
					port, err = strconv.Atoi(args[0])
					if err != nil {
						port = 8081
					}
				}

				addr := fmt.Sprintf(":%d", port)
				return http.ListenAndServe(addr, mux)
			})

			if err := g.Wait(); err != nil {
				cmd.PrintErrln("run worker", err)
			}

			return nil
		},
	}

	return cmd
}
