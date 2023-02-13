package ctxlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

type fieldsKey struct{}

type Fields map[string]any

type CtxLogger struct {
	Fields Fields
	Log    *logrus.Logger
}

func New(options ...func(*CtxLogger)) *CtxLogger {
	log := &CtxLogger{
		Log: logrus.New(),
	}

	for _, o := range options {
		o(log)
	}

	return log
}

func WithJSONFormat() func(*CtxLogger) {
	return func(l *CtxLogger) {
		l.Log.SetFormatter(&logrus.JSONFormatter{})
	}
}

func getFields(ctx context.Context) Fields {
	if val, ok := ctx.Value(fieldsKey{}).(Fields); ok {
		return val
	}
	return Fields{}
}

func WithFields(ctx context.Context, newFields ...Fields) context.Context {
	fields := Fields{}

	for k, v := range getFields(ctx) {
		fields[k] = v
	}

	for _, fieldMap := range newFields {
		for k, v := range fieldMap {
			fields[k] = v
		}
	}

	return context.WithValue(ctx, fieldsKey{}, fields)
}

func convertFieldsToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := make(logrus.Fields)
	for k, v := range fields {
		logrusFields[k] = v
	}
	return logrusFields
}
