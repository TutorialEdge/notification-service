package ctxlog

import (
	"context"
)

func (l *CtxLogger) Info(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Info(msg)
}
