package config

const (
	DefaultHttpAddr = "127.0.0.1"
	DefaultHttpPort = 4748
	DefaultRunMode  = ModeDebug

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
}

// NewServerDefaultConf returns a default configuration which is suitable for
// local/debug mode.
func NewServerDefaultConf() *Server {
	return &Server{
		HTTPAddr: DefaultHttpAddr,
		HTTPPort: DefaultHttpPort,
		RunMode:  DefaultRunMode,
	}
}

// NewServerConfig returns a Server Configuration based on params
func NewServerConfig(addr string, port uint, runMode string) *Server {
	return &Server{
		HTTPAddr: addr,
		HTTPPort: port,
		RunMode:  runMode,
	}
}
