package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Logger(logg *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logg.With(
			slog.Attr{
				Key:   "method",
				Value: slog.StringValue(c.Request.Method),
			},
			slog.Attr{
				Key:   "path",
				Value: slog.StringValue(c.Request.URL.Path),
			},
			slog.Attr{
				Key:   "query",
				Value: slog.StringValue(c.Request.URL.RawQuery),
			},
		)
		c.Next()
		logger.Info("request completed")
	}
}
