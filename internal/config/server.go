package config

import (
	"fmt"
	"time"
)

const (
	DefaultHttpAddr     = "127.0.0.1"
	DefaultHttpPort     = 4748
	DefaultRunMode      = ModeDebug
	DefaultReadTimeout  = time.Minute
	DefaultWriteTimeout = time.Minute

	ModeDebug   = "debug"
	ModeRelease = "release"
	ModeTest    = "test"
)

// Server configuration is being used by the API Web Server
type Server struct {
	HTTPAddr string
	HTTPPort uint

	// RunMode indicates the mode which the server should run in
	// Examples: debug, release, test
	RunMode string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewServerDefaultConf returns a default configuration which is suitable for
// local/debug mode.
func NewServerDefaultConf() *Server {
	return &Server{
		HTTPAddr:     DefaultHttpAddr,
		HTTPPort:     DefaultHttpPort,
		RunMode:      DefaultRunMode,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
	}
}

// NewServerConfig returns a Server Configuration based on params
func NewServerConfig(addr string, port uint, runMode string,
	readTimeout, writeTimeout time.Duration) *Server {

	return &Server{
		HTTPAddr:     addr,
		HTTPPort:     port,
		RunMode:      runMode,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

// AddrString returns a formatted ip:port string
// example: 127.0.0.1:8080
func (sc *Server) AddrString() string {
	return fmt.Sprintf("%s:%d", sc.HTTPAddr, sc.HTTPPort)
}
