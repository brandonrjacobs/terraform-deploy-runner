package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"tracks-grpc/internal"
)

func New(cfg *viper.Viper) (*zap.Logger, error) {
	zconfig := zapConfig(cfg)
	zlog, _ := zconfig.Build()
	return zlog, nil
}

func NewStartUp(cfg *viper.Viper) internal.StartupLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("startup")
	zWrapper := zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &zWrapper
}

func NewBackgroundLog(cfg *viper.Viper) internal.BackgroundLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("background")
	zWrapper := zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &background{zLog: &zWrapper}
}

func NewRequestLog(cfg *viper.Viper) internal.RequestLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("request")
	zWrapper := &zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &request{zLog: zWrapper}
}
