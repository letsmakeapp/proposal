package logger

import "context"

type Interface interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
}

type LoggerOptions struct {
	args []any
	ctx  context.Context
}

type LoggerOption func(opts *LoggerOptions)

func WithArgs(args ...any) LoggerOption {
	return func(opts *LoggerOptions) {
		opts.args = args
	}
}

func WithContext(ctx context.Context) LoggerOption {
	return func(opts *LoggerOptions) {
		opts.ctx = ctx
	}
}
