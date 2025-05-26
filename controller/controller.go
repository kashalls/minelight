package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kashalls/minelight/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

func Run() {
	cfg := internal.InitServerConfig()
	log.Infof("config: %s", cfg)

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Invalid log level: %s", cfg.LogLevel)
	}
	log.SetLevel(logLevel)

	if cfg.DryRun {
		log.Info("Running in dry-run mode, no changes will be applied.")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var g errgroup.Group

	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", cfg.MetricsHost, cfg.MetricsPort)
		http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		})

		log.Debugf("serving 'healthz' on '%s/healthz'", addr)
		log.Debugf("serving 'metrics' on '%s/metrics'", addr)
		log.Debugf("registered '%d' metrics", len(metrics.RegisterMetric.Metrics))

		http.Handle("/metrics", promhttp.Handler())

		log.Fatal(http.ListenAndServe(addr, nil))
		return nil
	})

	g.Go(func() error {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-sigCh:
			log.Infof("Received signal: %s, shutting down...", sig)
			cancel()
			return nil
		}
	})

	if cfg.Dashboard {
		g.Go(func() error {
			return nil
		})
	}

	if err := g.Wait(); err != nil && err != context.Canceled {
		log.Fatalf("Server exited with error: %v", err)
	}

}
