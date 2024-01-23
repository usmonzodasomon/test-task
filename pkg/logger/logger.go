package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Inf(inf string) slog.Attr {
	return slog.Attr{
		Key:   "info",
		Value: slog.StringValue(inf),
	}
}

func Deb(deb string) slog.Attr {
	return slog.Attr{
		Key:   "debug",
		Value: slog.StringValue(deb),
	}
}

func Warn(warn string) slog.Attr {
	return slog.Attr{
		Key:   "warning",
		Value: slog.StringValue(warn),
	}
}
