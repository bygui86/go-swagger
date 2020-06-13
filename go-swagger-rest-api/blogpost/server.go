package blogpost

import (
	"context"
	"time"

	"github.com/bygui86/go-swagger/go-swagger-rest-api/logging"
)

func NewRestServer() *Server {
	logging.Log.Debug("Create new REST server")

	logging.Log.Debug("Load configurations")
	cfg := loadConfig()

	logging.Log.Debug("Create router")
	router := setupRouter()

	logging.Log.Debug("Create HTTP server")
	httpServer := setupHttpServer(router, cfg.RestHost, cfg.RestPort)

	logging.Log.Debug("Create REST server")
	server := &Server{
		config:     cfg,
		router:     router,
		httpServer: httpServer,
		errChannel: make(chan error, 1),
		blogPosts:  initBlogPosts(),
		routes:     initRoutes(),
	}

	logging.Log.Debug("Setup handlers")
	server.setupHandlers()

	return server
}

func (s *Server) Start() {
	logging.Log.Info("Start REST server")

	if s.router != nil && s.httpServer != nil && !s.running {
		go s.startHttpServerController()

		go s.listenAndServe()

		s.running = true
		logging.SugaredLog.Infof("REST server listening on port %d", s.config.RestPort)
		return
	}

	logging.Log.Error("REST server start failed: router, HTTP server not initialized or HTTP server already running")
}

func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown REST server, timeout %d", timeout)

	if s.router != nil && s.httpServer != nil && s.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("REST server shutdown failed: %s", err.Error())
		}
		s.running = false
		return
	}

	logging.Log.Error("REST server shutdown failed: router, HTTP server not initialized or HTTP server not running")
}
