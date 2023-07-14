package logger

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InitLogger(logFormat string) {
	switch strings.ToLower(logFormat) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "gce":
		//
	default:
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	}
}

func Info(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Infof(format, args...)
}

func Warn(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Warnf(format, args...)
}

func Debug(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Debugf(format, args...)
}

func Error(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Errorf(format, args...)
}

func Fatal(ctx context.Context, ctxName string, format string, args ...interface{}) {
	getEntry(ctx, ctxName).Fatalf(format, args...)
}

func getEntry(ctx context.Context, ctxName string) *log.Entry {
	return log.WithFields(log.Fields{
		"context": ctxName,
		// "correlationId": ctx.Value(utils.CORRELATION_ID),
	})
}
