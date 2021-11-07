package logging

import (
	"tracks-grpc/config"
	"tracks-grpc/internal"
)

var Component = internal.NewComponent(
	"logging",
	[]config.EnvVar{
		config.EnvLogFormat,
		config.EnvLogLevel,
		config.EnvLogSampleEvery,
		config.EnvLogSampleInitial,
		config.EnvLogSamplingRate},
	NewBackgroundLog,
	NewRequestLog,
)
