package logging

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"tracks-grpc/config"
)

func TestZapLog(t *testing.T) {
	t.Run("TestInfo", testInfo)
}

func testInfo(t *testing.T) {
	cfg := viper.New()
	cfg.Set(config.LogFormat.String(), "json")
	cfg.Set(config.LogLevel.String(), "info")
	cfg.Set(config.DevMode.String(), true)
	z, err := New(cfg)
	z = z.Named("test")
	z = z.With(zap.String("test_Val", "test1"))
	assert.NoError(t, err)
	log := zapLog{
		logger:  z,
		sugared: z.Sugar(),
	}

	log.Infof("template: %s", "val1")

}
