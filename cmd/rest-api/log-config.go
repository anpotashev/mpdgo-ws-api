package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/anpotashev/mpd-ws-api/internal/api/middleware"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

func configLogs(config *AppConfig) {

	//logHandler := slog.NewJSONHandler(
	//	os.Stdout,
	//	&slog.HandlerOptions{
	//		Level: slog.LevelDebug,
	//	})
	logHandler := &contextHandler{slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: getLogLevel(config.LogLevel),
		})}
	//tintLogHandler := tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen})
	//logHandler := tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen})
	log := slog.New(logHandler)
	slog.SetDefault(log)
	mpdapi.SetLogger(log)
	//liblog := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}))
	//mpdapi.SetLogger(liblog)
}

type contextHandler struct {
	slog.Handler
}

func (contextHandler *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if rId, ok := ctx.Value(middleware.RequestIdContextAttributeName).(string); ok && rId != "" {
		r.AddAttrs(slog.String("request_id", rId))
	}
	return contextHandler.Handler.Handle(ctx, r)
}
