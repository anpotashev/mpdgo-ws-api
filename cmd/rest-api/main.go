package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/anpotashev/mpd-ws-api/internal/api/middleware"
	v1 "github.com/anpotashev/mpd-ws-api/internal/api/v1"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}
	configLogs(config)
	api, err := mpdapi.NewMpdApi(ctx,
		config.MpdHost,
		config.MpdPort,
		config.MpdPassword,
		true,
		config.MaxBatchCommand,
		config.MpdPoolSize,
		time.Millisecond*time.Duration(config.CommandReadIntervalMillis),
		time.Second*time.Duration(config.PingIntervalSeconds),
	)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(middleware.LoggerContextMiddleware)
	v1.New(router.PathPrefix("/v1").Subrouter(), api)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	panic(srv.ListenAndServe())
}

func getLogLevel(level string) slog.Leveler {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic("Invalid log level")
	}
}
