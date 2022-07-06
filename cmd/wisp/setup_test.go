package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/joatisio/wisp/internal/config"
	"github.com/stretchr/testify/assert"
)

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
			got := generateDsn(&config.Config{Database: tt.dbConf})

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
	conf := &config.Config{
		Server: config.Server{
			HTTPAddr: "0.0.0.0",
			HTTPPort: 8080,
			RunMode:  config.ModeDebug,
		},
		Database: config.Database{
			Username: "root",
			Password: "root",
			Host:     "127.0.0.1",
			DBName:   "joatis",
			Port:     5432,
		},
	}
	pre := ConfigPrefixEnv + "_"
	os.Setenv(pre+"HTTP_ADDR", conf.Server.HTTPAddr)
	os.Setenv(pre+"HTTP_PORT", strconv.FormatUint(uint64(conf.Server.HTTPPort), 10))
	os.Setenv(pre+"HTTP_MODE", conf.Server.RunMode)
	os.Setenv(pre+"DB_USERNAME", conf.Database.Username)
	os.Setenv(pre+"DB_PASSWORD", conf.Database.Password)
	os.Setenv(pre+"DB_HOST", conf.Database.Host)
	os.Setenv(pre+"DB_NAME", conf.Database.DBName)
	os.Setenv(pre+"DB_PORT", strconv.FormatUint(uint64(conf.Database.Port), 10))

	c := setupConfig()
	assert.EqualValues(t, conf, c)
}
