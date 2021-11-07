package logging

import (
	"context"
	"go.uber.org/zap"
	"tracks-grpc/internal"
)

type background struct {
	zLog *zapLog
}

func (b background) Debug(args ...interface{}) {
	b.zLog.Debug(args...)
}

func (b background) Debugf(template string, args ...interface{}) {
	b.zLog.Debugf(template, args...)
}

func (b background) Debugw(message string, keyValues ...interface{}) {
	b.zLog.Debugw(message, keyValues...)
}

func (b background) Info(args ...interface{}) {
	b.zLog.Info(args...)
}

func (b background) Infof(template string, args ...interface{}) {
	b.zLog.Infof(template, args...)
}

func (b background) Infow(message string, keyValues ...interface{}) {
	b.zLog.Infow(message, keyValues...)
}

func (b background) Warn(args ...interface{}) {
	b.zLog.Warn(args...)
}

func (b background) Warnf(template string, args ...interface{}) {
	b.zLog.Warnf(template, args...)
}

func (b background) Warnw(message string, keyValues ...interface{}) {
	b.zLog.Warnw(message, keyValues...)
}

func (b background) Error(args ...interface{}) {
	b.zLog.Error(args...)
}

func (b background) Errorf(template string, args ...interface{}) {
	b.zLog.Errorf(template, args...)
}

func (b background) Errorw(message string, keyValues ...interface{}) {
	b.zLog.Errorw(message, keyValues...)
}

func (b background) Err(err error, message string) {
	b.zLog.Err(err, message)
}

func (b background) Errf(err error, template string, args ...interface{}) {
	b.zLog.Errf(err, template, args...)
}

func (b background) Errw(err error, message string, keyValues ...interface{}) {
	b.zLog.Errw(err, message, keyValues...)
}

func (b background) DebugCtx(ctx context.Context, message string) {
	b.zLog.DebugCtx(ctx, message)
}

func (b background) DebugfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.DebugfCtx(ctx, template, args...)
}

func (b background) DebugwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.DebugwCtx(ctx, message, keyValues...)
}

func (b background) InfoCtx(ctx context.Context, message string) {
	b.zLog.InfoCtx(ctx, message)
}

func (b background) InfofCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.InfofCtx(ctx, template, args...)
}

func (b background) InfowCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.InfowCtx(ctx, message, keyValues...)
}

func (b background) WarnCtx(ctx context.Context, message string) {
	b.zLog.WarnCtx(ctx, message)
}

func (b background) WarnfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.WarnfCtx(ctx, template, args...)
}

func (b background) WarnwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.WarnwCtx(ctx, message, keyValues...)
}

func (b background) ErrorCtx(ctx context.Context, message string) {
	b.zLog.ErrorCtx(ctx, message)
}

func (b background) ErrorfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.ErrorfCtx(ctx, template, args...)
}

func (b background) ErrorwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.ErrorwCtx(ctx, message, keyValues...)
}

func (b background) ErrCtx(ctx context.Context, err error, message string) {
	b.zLog.ErrCtx(ctx, err, message)
}

func (b background) ErrfCtx(ctx context.Context, err error, template string, args ...interface{}) {
	b.zLog.ErrfCtx(ctx, err, template, args...)
}

func (b background) ErrwCtx(ctx context.Context, err error, message string, keyValues ...interface{}) {
	b.zLog.ErrwCtx(ctx, err, message, keyValues...)
}

func (b background) ChildLog(name string) internal.BackgroundLog {
	newZ := b.zLog.logger.With(zap.String("name", name))
	return &background{zLog: &zapLog{logger: newZ, sugared: newZ.Sugar()}}
}
