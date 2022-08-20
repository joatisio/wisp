package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joatisio/wisp/internal/config"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

const maxHeaderBytes = 1 << 20

type IRouter interface {
	SetupRoutes(*gin.RouterGroup)
}

type Server struct {
	logger      *zap.Logger
	routers     []IRouter
	engine      *gin.Engine
	config      *config.Server
	routerGroup *gin.RouterGroup
	apiPrefix   string
	mu          sync.Mutex
}

func NewServer(engine *gin.Engine, logger *zap.Logger, c *config.Server, apiPrefix string) *Server {
	return &Server{
		engine:      engine,
		logger:      logger,
		config:      c,
		apiPrefix:   apiPrefix,
		routerGroup: engine.Group(apiPrefix),
	}
}

func (s *Server) AddRouter(r IRouter) *Server {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.routers = append(s.routers, r)
	return s
}

func (s *Server) AddRouters(rs ...IRouter) *Server {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.routers = append(s.routers, rs...)
	return s
}

func (s *Server) RegisterRoutes() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, rtr := range s.routers {
		rtr.SetupRoutes(s.routerGroup)
	}
}

// Serve always returns a non-nil error.
func (s *Server) Serve(addr string) error {
	srv := &http.Server{
		Addr:           addr,
		Handler:        s.engine,
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	s.logger.Info(fmt.Sprintf("server is running on %s", addr))

	if err := srv.ListenAndServe(); err != nil { // It is always non-nil, but looks like linter has problems with it
		return fmt.Errorf("failed to launch server: %w", err)
	}

	return nil
}
