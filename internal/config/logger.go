package config

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

type Logger struct {
	Level             LogLevel
	DisableCaller     bool
	DisableStacktrace bool
}

func NewDevelopmentConfig() *Logger {
	return &Logger{
		Level:             LevelDebug,
		DisableCaller:     false,
		DisableStacktrace: false,
	}
}

func NewProductionConfig() *Logger {
	return &Logger{
		Level:             LevelError,
		DisableCaller:     false,
		DisableStacktrace: true,
	}
}

func NewLoggerConfig(lvl LogLevel, disableCaller bool, disableStack bool) *Logger {
	return &Logger{
		Level:             lvl,
		DisableCaller:     disableCaller,
		DisableStacktrace: disableStack,
	}
}
