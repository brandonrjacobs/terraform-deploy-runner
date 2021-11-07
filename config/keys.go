package config

type Key string

func (k Key) String() string {
	return string(k)
}

var AppEnv Key = "APP_ENV"
var LogFormat Key = "LOG_FORMAT"
var LogLevel Key = "LOG_LEVEL"
var DevMode Key = "DEV_MODE"
var LogSampleEvery Key = "LOG_SAMPLE_EVERY"
var LogSampleInitial Key = "LOG_SAMPLE_INITIAL"
var LogSamplingRate Key = "LOG_SAMPLING_RATE"

// AllowedGitRepositories This key represents a struct of repository url -> key for pulling the repository
var AllowedGitRepositories Key = "ALLOWED_GIT_REPOSITORIES"


