package logging

import "context"

type log struct {
	zLog *zapLog
	ctx  context.Context
}

func (l log) Debug(args ...interface{}) {
	l.zLog.Debug(args...)
}

func (l log) Debugf(template string, args ...interface{}) {
	l.zLog.DebugfCtx(l.ctx, template, args...)
}

func (l log) Debugw(message string, keyValues ...interface{}) {
	l.zLog.DebugwCtx(l.ctx, message, keyValues...)
}

func (l log) Info(args ...interface{}) {
	l.zLog.Info(args...)
}

func (l log) Infof(template string, args ...interface{}) {
	l.zLog.InfofCtx(l.ctx, template, args...)
}

func (l log) Infow(message string, keyValues ...interface{}) {
	l.zLog.InfowCtx(l.ctx, message, keyValues...)
}

func (l log) Warn(args ...interface{}) {
	l.zLog.Warn(args...)
}

func (l log) Warnf(template string, args ...interface{}) {
	l.zLog.WarnfCtx(l.ctx, template, args...)
}

func (l log) Warnw(message string, keyValues ...interface{}) {
	l.zLog.WarnwCtx(l.ctx, message, keyValues...)
}

func (l log) Error(args ...interface{}) {
	l.zLog.Error(args...)
}

func (l log) Errorf(template string, args ...interface{}) {
	l.zLog.ErrorfCtx(l.ctx, template, args...)
}

func (l log) Errorw(message string, keyValues ...interface{}) {
	l.zLog.ErrorwCtx(l.ctx, message, keyValues...)
}

func (l log) Err(err error, message string) {
	l.zLog.ErrCtx(l.ctx, err, message)
}

func (l log) Errf(err error, template string, args ...interface{}) {
	l.zLog.ErrfCtx(l.ctx, err, template, args...)
}

func (l log) Errw(err error, message string, keyValues ...interface{}) {
	l.zLog.ErrwCtx(l.ctx, err, message, keyValues...)
}
