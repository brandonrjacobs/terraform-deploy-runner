package config

type EnvVar struct {
	Key         Key    `json:"key"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

var EnvLogFormat = EnvVar{
	Key:         LogFormat,
	Name:        "LOG_FORMAT",
	Description: "Log format",
}

var EnvLogLevel = EnvVar{
	Key:         LogLevel,
	Name:        "LOG_LEVEL",
	Description: "Log level",
}

var EnvLogSamplingRate = EnvVar{
	Key:         LogSamplingRate,
	Name:        "LOG_SAMPLING_RATE",
	Description: "Sampling rate for logging",
}

var EnvDevMode = EnvVar{
	Key:         DevMode,
	Name:        "DEV_MODE",
	Description: "Development mode boolean",
}

var EnvLogSampleEvery = EnvVar{
	Key:         LogSampleEvery,
	Name:        "LOG_SAMPLE_EVERY",
	Description: "Log sampling every request",
}

var EnvLogSampleInitial = EnvVar{
	Key:         LogSampleInitial,
	Name:        "LOG_SAMPLE_INITIAL",
	Description: "Log sampling initial request",
}

var EnvGrpcAddress = EnvVar{
	Key:         GrpcAddress,
	Name:        "GRPC_ADDRESS",
	Description: "The address port of the grpc server",
}

var EnvKafkaBootstrapServerAddress = EnvVar{
	Key:         KafkaBootstrapServerAddress,
	Name:        "KAFKA_BOOTSTRAP_SERVER_ADDRESS",
	Description: "",
}
