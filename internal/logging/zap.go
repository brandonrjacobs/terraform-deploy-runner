package logging

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"tracks-grpc/config"
)

var samplingRate int
var sampleEvery int

func zapConfig(cfg *viper.Viper) zap.Config {
	// Determine logger format
	encoding := strings.ToLower(cfg.GetString(config.LogFormat.String()))
	var encoderConfig zapcore.EncoderConfig
	switch encoding {
	case "json":
		encoderConfig = zap.NewProductionEncoderConfig()
	case "console":
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	default:
		fmt.Printf("Unknown log format: '%v'\n", encoding)
		os.Exit(1)
	}

	var level zap.AtomicLevel
	levelString := strings.ToLower(cfg.GetString(config.LogLevel.String()))
	switch levelString {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		fmt.Printf("Unknown log level: '%v'\n", levelString)
		os.Exit(1)
	}

	var sampling *zap.SamplingConfig
	samplingRate = cfg.GetInt(config.LogSamplingRate.String())
	if samplingRate > 0 {
		sampleEvery = cfg.GetInt(config.LogSampleEvery.String())
		if sampleEvery <= 0 {
			sampleEvery = 50
		}

		sampling = &zap.SamplingConfig{
			Initial:    samplingRate,
			Thereafter: sampleEvery,
		}
	}

	return zap.Config{
		Development:      cfg.GetBool(config.DevMode.String()),
		Encoding:         encoding,
		EncoderConfig:    encoderConfig,
		Level:            level,
		Sampling:         sampling,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

type zapLog struct {
	logger  *zap.Logger
	sugared *zap.SugaredLogger
}

func (w *zapLog) Debug(args ...interface{}) {
	w.sugared.Debug(args...)
}

func (w *zapLog) Debugf(template string, args ...interface{}) {
	w.sugared.Debugf(template, args...)
}

func (w *zapLog) Debugw(message string, keyValues ...interface{}) {
	w.sugared.Debugw(message, keyValues...)
}

func (w *zapLog) DebugCtx(ctx context.Context, message string) {
	w.logger.Debug(message, getContextFields(ctx)...)
}

func (w *zapLog) DebugfCtx(ctx context.Context, template string, args ...interface{}) {
	w.logger.Debug(fmt.Sprintf(template, args...), getContextFields(ctx)...)
}

func (w *zapLog) DebugwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	w.sugared.Debugw(message, appendContextFields(ctx, keyValues...)...)
}

func (w *zapLog) Info(args ...interface{}) {
	w.sugared.Info(args...)
}

func (w *zapLog) Infof(template string, args ...interface{}) {
	w.sugared.Infof(template, args...)
}

func (w *zapLog) Infow(message string, keyValues ...interface{}) {
	w.sugared.Infow(message, keyValues...)
}

func (w *zapLog) InfoCtx(ctx context.Context, message string) {
	w.logger.Info(message, getContextFields(ctx)...)
}

func (w *zapLog) InfofCtx(ctx context.Context, template string, args ...interface{}) {
	w.logger.Info(fmt.Sprintf(template, args...), getContextFields(ctx)...)
}

func (w *zapLog) InfowCtx(ctx context.Context, message string, keyValues ...interface{}) {
	w.sugared.Infow(message, appendContextFields(ctx, keyValues...)...)
}

func (w *zapLog) Warn(args ...interface{}) {
	w.sugared.Warn(args...)
}

func (w *zapLog) Warnf(template string, args ...interface{}) {
	w.sugared.Warnf(template, args...)
}

func (w *zapLog) Warnw(message string, keyValues ...interface{}) {
	w.sugared.Warnw(message, keyValues...)
}

func (w *zapLog) WarnCtx(ctx context.Context, message string) {
	w.logger.Warn(message, getContextFields(ctx)...)
}

func (w *zapLog) WarnfCtx(ctx context.Context, template string, args ...interface{}) {
	w.logger.Warn(fmt.Sprintf(template, args...), getContextFields(ctx)...)
}

func (w *zapLog) WarnwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	w.sugared.Warnw(message, appendContextFields(ctx, keyValues...)...)
}

func (w *zapLog) Error(args ...interface{}) {
	w.sugared.Error(args...)
}

func (w *zapLog) Errorf(template string, args ...interface{}) {
	w.sugared.Errorf(template, args...)
}

func (w *zapLog) Errorw(message string, keyValues ...interface{}) {
	w.sugared.Errorw(message, keyValues...)
}

func (w *zapLog) ErrorCtx(ctx context.Context, message string) {
	w.logger.Error(message, getContextFields(ctx)...)
}

func (w *zapLog) ErrorfCtx(ctx context.Context, template string, args ...interface{}) {
	w.logger.Error(fmt.Sprintf(template, args...), getContextFields(ctx)...)
}

func (w *zapLog) ErrorwCtx(ctx context.Context, message string, keyValues ...interface{}) {
	w.sugared.Errorw(message, appendContextFields(ctx, keyValues...)...)
}

func (w *zapLog) Err(err error, message string) {
	w.logger.Error(message, zap.String(FieldError.String(), err.Error()))
}

func (w *zapLog) Errf(err error, template string, args ...interface{}) {
	w.logger.Error(fmt.Sprintf(template, args...), zap.String(FieldError.String(), err.Error()))
}

func (w *zapLog) Errw(err error, message string, keyValues ...interface{}) {
	fields := make([]interface{}, 0, len(keyValues)+2)
	fields = append(fields, FieldError, err.Error())
	fields = append(fields, keyValues...)
	w.sugared.Errorw(message, fields...)
}

func (w *zapLog) ErrCtx(ctx context.Context, err error, message string) {
	w.logger.Error(message, zap.String(FieldError.String(), err.Error()), zap.String(FieldRequestID.String(), FieldRequestID.GetFromContext(ctx)))
}

func (w *zapLog) ErrfCtx(ctx context.Context, err error, template string, args ...interface{}) {
	w.logger.Error(fmt.Sprintf(template, args...), zap.String(FieldError.String(), err.Error()), zap.String(FieldRequestID.String(), FieldRequestID.GetFromContext(ctx)))
}

func (w *zapLog) ErrwCtx(ctx context.Context, err error, message string, keyValues ...interface{}) {
	fields := make([]interface{}, 0, len(keyValues)+6)
	fields = append(fields, FieldError.String(), err.Error(), FieldRequestID.String(), FieldRequestID.GetFromContext(ctx))
	w.sugared.Errorw(message, fields...)
}

func getContextFields(ctx context.Context) []zap.Field {
	return []zap.Field{zap.String(FieldRequestID.String(), FieldRequestID.GetFromContext(ctx))}
}

func appendContextFields(ctx context.Context, fieldsAndValues ...interface{}) []interface{} {
	results := make([]interface{}, 0, len(fieldsAndValues)+4)
	results = append(results, fieldsAndValues...)
	return append(results, FieldRequestID.String(), FieldRequestID.GetFromContext(ctx))
}
