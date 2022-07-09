package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/joatisio/wisp/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const ConfigPrefixEnv = "WISP"

// initConfig reads configuration from environment variables and
// returns a Config object
func setupConfig() *config.Config {
	viper.SetEnvPrefix(ConfigPrefixEnv)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &config.Config{
		Server: &config.Server{
			HTTPAddr: viper.GetString("http.addr"),
			HTTPPort: viper.GetUint("http.port"),
			RunMode:  viper.GetString("http.mode"),
		},
		Database: &config.Database{
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			DBName:   viper.GetString("db.name"),
			Port:     viper.GetUint("db.port"),
			Secure:   viper.GetBool("db.secure"),
		},
		Logger: &config.Logger{
			Level:             config.LogLevel(viper.GetString("logger.level")),
			DisableCaller:     viper.GetBool("logger.disablecaller"),
			DisableStacktrace: viper.GetBool("logger.disablestacktrace"),
		},
	}
}

func generateDsn(c *config.Database) string {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	return fmt.Sprintf(dsn, c.Host, c.Username, c.Password, c.DBName, c.Port)
}

// createDialector returns a postgres dialector
func createDialector(dsn string, conn *sql.DB) gorm.Dialector {
	var con *sql.DB
	if conn != nil {
		con = conn
	} else {
		c, err := sql.Open("postgres", dsn)
		if err != nil {
			panic("cannot open database connection")
		}
		con = c
	}

	return postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  dsn,
		PreferSimpleProtocol: true,
		WithoutReturning:     false,
		Conn:                 con,
	})
}

func setupDatabase(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// toZapLevel accepts a level param and returns a zap compatible
// level respectively. It returns Error Level if there is no match.
func toZapLevel(level config.LogLevel) zapcore.Level {
	switch level {
	case config.LevelDebug:
		return zap.DebugLevel
	case config.LevelInfo:
		return zap.InfoLevel
	case config.LevelWarn:
		return zap.WarnLevel
	case config.LevelError:
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}

// setupLogger returns a zap logger based on Logger configuration
func setupLogger(c *config.Logger) *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(toZapLevel(c.Level)),
		Development:       false,
		DisableCaller:     c.DisableCaller,
		DisableStacktrace: c.DisableStacktrace,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			MessageKey:     "message",
			TimeKey:        "time",
			NameKey:        "name",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		InitialFields:    nil,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
