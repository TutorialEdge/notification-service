package ctxlog

import (
	"context"
)

func (l *CtxLogger) Warn(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Warn(msg)
}
