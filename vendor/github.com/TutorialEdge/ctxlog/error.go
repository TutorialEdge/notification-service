package ctxlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (l *CtxLogger) Error(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Error(msg)
}

func (l *CtxLogger) ErrorFn(ctx context.Context, logFunc logrus.LogFunction) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Logger.ErrorFn(logFunc)
}
