package logging

import (
	"context"
	"tracks-grpc/internal"
)

type request struct {
	zLog *zapLog
}

func (b request) ForContext(ctx context.Context) internal.Log {
	return &log{
		zLog: b.zLog,
		ctx:  ctx,
	}
}

func (b request) Debug(args ...interface{}) {
	b.zLog.Debug(args...)
}

func (b request) Debugf(template string, args ...interface{}) {
	b.zLog.Debugf(template, args...)
}

func (b request) Debugw(message string, keyValues ...interface{}) {
	b.zLog.Debugw(message, keyValues...)
}

func (b request) Info(args ...interface{}) {
	b.zLog.Info(args...)
}

func (b request) Infof(template string, args ...interface{}) {
	b.zLog.Infof(template, args...)
}

func (b request) Infow(message string, keyValues ...interface{}) {
	b.zLog.Infow(message, keyValues...)
}

func (b request) Warn(args ...interface{}) {
	b.zLog.Warn(args...)
}

func (b request) Warnf(template string, args ...interface{}) {
	b.zLog.Warnf(template, args...)
}

func (b request) Warnw(message string, keyValues ...interface{}) {
	b.zLog.Warnw(message, keyValues...)
}

func (b request) Error(args ...interface{}) {
	b.zLog.Error(args...)
}

func (b request) Errorf(template string, args ...interface{}) {
	b.zLog.Errorf(template, args...)
}

func (b request) Errorw(message string, keyValues ...interface{}) {
	b.zLog.Errorw(message, keyValues...)
}

func (b request) Err(err error, message string) {
	b.zLog.Err(err, message)
}

func (b request) Errf(err error, template string, args ...interface{}) {
	b.zLog.Errf(err, template, args...)
}

func (b request) Errw(err error, message string, keyValues ...interface{}) {
	b.zLog.Errw(err, message, keyValues...)
}

func (b request) DebugCtx(ctx context.Context, message string) {
	b.zLog.DebugCtx(ctx, message)
}

func (b request) DebugfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.DebugfCtx(ctx, template, args...)
}

func (b request) DebugwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.DebugwCtx(ctx, message, keyValues...)
}

func (b request) InfoCtx(ctx context.Context, message string) {
	b.zLog.InfoCtx(ctx, message)
}

func (b request) InfofCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.InfofCtx(ctx, template, args...)
}

func (b request) InfowCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.InfowCtx(ctx, message, keyValues...)
}

func (b request) WarnCtx(ctx context.Context, message string) {
	b.zLog.WarnCtx(ctx, message)
}

func (b request) WarnfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.WarnfCtx(ctx, template, args...)
}

func (b request) WarnwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.WarnwCtx(ctx, message, keyValues...)
}

func (b request) ErrorCtx(ctx context.Context, message string) {
	b.zLog.ErrorCtx(ctx, message)
}

func (b request) ErrorfCtx(ctx context.Context, template string, args ...interface{}) {
	b.zLog.ErrorfCtx(ctx, template, args...)
}

func (b request) ErrorwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	b.zLog.ErrorwCtx(ctx, message, keyValues...)
}

func (b request) ErrCtx(ctx context.Context, err error, message string) {
	b.zLog.ErrCtx(ctx, err, message)
}

func (b request) ErrfCtx(ctx context.Context, err error, template string, args ...interface{}) {
	b.zLog.ErrfCtx(ctx, err, template, args...)
}

func (b request) ErrwCtx(ctx context.Context, err error, message string, keyValues ...interface{}) {
	b.zLog.ErrwCtx(ctx, err, message, keyValues...)
}
