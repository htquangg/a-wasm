package constants

import "time"

const (
	AppEnv     = "APP_ENV"
	ConfigPath = "CONFIG_PATH"
	LogLevel   = "LOG_LEVEL"
	LogPath    = "LOG_PATH"

	Yaml = "yaml"

	Dev        = "development"
	Test       = "test"
	Production = "production"

	ReadTimeout       = 10 * time.Minute
	WriteTimeout      = 60 * time.Second
	ReadHeaderTimeout = 30 * time.Second
	ShutdownTimeout   = 5 * time.Second

	TOKEN_LENGTH = 32
)
