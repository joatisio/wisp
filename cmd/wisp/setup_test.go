package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/joatisio/wisp/internal/config"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var testConf = &config.Config{
	Server: &config.Server{
		HTTPAddr: "0.0.0.0",
		HTTPPort: 8080,
		RunMode:  config.ModeDebug,
	},
	Database: &config.Database{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		DBName:   "joatis",
		Port:     5432,
		Secure:   false,
	},
	Logger: &config.Logger{
		Level:             config.LogLevel(config.LevelDebug),
		DisableCaller:     false,
		DisableStacktrace: false,
	},
}

func Test_generateDsn(t *testing.T) {

	tests := []struct {
		name      string
		dbConf    config.Database
		want      string
		isInvalid bool
	}{
		{
			name: "valid",
			dbConf: config.Database{
				Username: "root",
				Password: "root",
				Host:     "127.0.0.1",
				DBName:   "joatis",
				Port:     5432,
				Secure:   false,
			},
			want:      "host=127.0.0.1 user=root password=root dbname=joatis port=5432 sslmode=disable",
			isInvalid: false,
		},
		{
			name: "invalid",
			dbConf: config.Database{
				Username: "root",
				Password: "root",
				Host:     "127.0.0.1",
				DBName:   "joatis",
				Port:     5432,
			},
			want:      "host=127.0.0.2 user=root password=root dbname=joatis port=5432 sslmode=disable",
			isInvalid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateDsn(&tt.dbConf)

			if !tt.isInvalid {
				if got != tt.want {
					t.Errorf("name: %v | generateDsn() got = %v, want %v", tt.name, got, tt.want)
				}
			} else {
				if got == tt.want {
					t.Errorf("name: %v | generateDsn() got a valid result, expected an invalid one", tt.name)
				}
			}
		})
	}
}

func Test_setupConfig(t *testing.T) {
	pre := ConfigPrefixEnv + "_"
	os.Setenv(pre+"HTTP_ADDR", testConf.Server.HTTPAddr)
	os.Setenv(pre+"HTTP_PORT", strconv.FormatUint(uint64(testConf.Server.HTTPPort), 10))
	os.Setenv(pre+"HTTP_MODE", testConf.Server.RunMode)

	os.Setenv(pre+"DB_USERNAME", testConf.Database.Username)
	os.Setenv(pre+"DB_PASSWORD", testConf.Database.Password)
	os.Setenv(pre+"DB_HOST", testConf.Database.Host)
	os.Setenv(pre+"DB_NAME", testConf.Database.DBName)
	os.Setenv(pre+"DB_SECURE", strconv.FormatBool(testConf.Database.Secure))
	os.Setenv(pre+"DB_PORT", strconv.FormatUint(uint64(testConf.Database.Port), 10))

	os.Setenv(pre+"LOGGER_LEVEL", string(testConf.Logger.Level))
	os.Setenv(pre+"LOGGER_DISABLECALLER", strconv.FormatBool(testConf.Logger.DisableCaller))
	os.Setenv(pre+"LOGGER_DISABLESTACKTRACE", strconv.FormatBool(testConf.Logger.DisableStacktrace))

	c := setupConfig()
	assert.EqualValues(t, testConf, c)
}

func Test_setupDatabase(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	dsn := generateDsn(testConf.Database)
	dialector := createDialector(dsn, db)
	gormDb := setupDatabase(dialector)
	ddb, err := gormDb.DB()
	if err != nil {
		t.Fatalf("cannot get DB object | error: %v", err)
	}
	assert.EqualValues(t, db, ddb)
}

func Test_setupLogger(t *testing.T) {
	// we should pass a config and check whether the passed level
	// is enabled. That's the only way that I know for testing setupLogger()
	conf := config.Logger{
		Level:             config.LogLevel(config.LevelError),
		DisableCaller:     false,
		DisableStacktrace: false,
	}

	z := setupLogger(&conf)
	lvl := toZapLevel(config.LogLevel(config.LevelDebug))
	assert.False(t, z.Core().Enabled(lvl))
}
