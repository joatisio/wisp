package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"

	"github.com/joatisio/wisp/internal/config"
	"github.com/spf13/viper"
)

const ConfigPrefixEnv = "WISP"

// initConfig will be executed on Cobra initiation
func setupConfig() *config.Config {
	viper.SetEnvPrefix(ConfigPrefixEnv)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &config.Config{
		Server: config.Server{
			HTTPAddr: viper.GetString("http.addr"),
			HTTPPort: viper.GetUint("http.port"),
			RunMode:  viper.GetString("http.mode"),
		},
		Database: config.Database{
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			DBName:   viper.GetString("db.name"),
			Port:     viper.GetUint("db.port"),
		},
	}
}

func generateDsn(c *config.Config) string {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	return fmt.Sprintf(dsn, c.Database.Host, c.Database.Username, c.Database.Password, c.Database.DBName, c.Database.Port)
}

func setupDatabase(c *config.Config) *gorm.DB {
	dsn := generateDsn(c)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// TODO we should read level from config
func setupLogger() *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
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
