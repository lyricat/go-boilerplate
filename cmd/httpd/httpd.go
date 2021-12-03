package httpd

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fox-one/pkg/logger"
	"github.com/fox-one/pkg/store/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lyricat/go-boilerplate/config"
	"github.com/lyricat/go-boilerplate/handler"
	"github.com/lyricat/go-boilerplate/handler/hc"
	"github.com/lyricat/go-boilerplate/session"
	"github.com/lyricat/go-boilerplate/store/wallet"
	"github.com/rs/cors"

	"github.com/drone/signal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCmdHttpd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "httpd [port]",
		Short: "start the httpd daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ctx := cmd.Context()
			s := session.From(ctx)

			database := db.MustOpen(config.C().DB)
			defer database.Close()

			client, err := s.GetClient()
			if err != nil {
				return err
			}

			pin, err := s.GetPin()
			if err != nil {
				return err
			}

			wallets := wallet.New(database, client, wallet.Config{Pin: pin})

			mux := chi.NewMux()
			mux.Use(middleware.Recoverer)
			mux.Use(middleware.StripSlashes)
			mux.Use(cors.AllowAll().Handler)
			mux.Use(logger.WithRequestID)
			mux.Use(middleware.Logger)
			mux.Use(middleware.NewCompressor(5).Handler)

			// /
			{
				mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("hello world"))
				})
			}

			// hc
			{
				mux.Mount("/hc", hc.Handle(cmd.Version))
			}

			// rpc & api
			{
				cfg := handler.Config{}
				svr := handler.New(cfg, wallets)
				// api v1
				restHandler := svr.HandleRest()
				mux.Mount("/api", restHandler)
			}

			port := 8080
			if len(args) > 0 {
				port, err = strconv.Atoi(args[0])
				if err != nil {
					port = 8080
				}
			}

			// launch server
			if err != nil {
				panic(err)
			}
			addr := fmt.Sprintf(":%d", port)

			svr := &http.Server{
				Addr:    addr,
				Handler: mux,
			}

			done := make(chan struct{}, 1)
			ctx = signal.WithContextFunc(ctx, func() {
				logrus.Debug("shutdown server...")

				// create context with timeout
				ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
				defer cancel()

				if err := svr.Shutdown(ctx); err != nil {
					logrus.WithError(err).Error("graceful shutdown server failed")
				}

				close(done)
			})

			logrus.Infoln("serve at", addr)
			if err := svr.ListenAndServe(); err != http.ErrServerClosed {
				logrus.WithError(err).Fatal("server aborted")
			}

			<-done
			return nil
		},
	}

	return cmd
}
