package http

import (
	"fmt"
	"net/http"
	"simple-job-board/pkg/config"
)

const (
	NilConfigError = "config is nil"
	NilRouterError = "router is nil"
)

type HttpServer struct {
	cfg    *config.Configuration
	router *http.ServeMux
}

func NewHttpServer(config *config.Configuration) *HttpServer {
	return &HttpServer{
		cfg:    config,
		router: http.NewServeMux(),
	}
}

func (s *HttpServer) addressBuilder() string {
	return s.cfg.HTTPServerHost + ":" + s.cfg.HTTPServerPort
}

func (s *HttpServer) Start() error {
	if s.cfg == nil {
		return fmt.Errorf(NilConfigError)
	}
	if s.router == nil {
		return fmt.Errorf(NilRouterError)
	}

	http.ListenAndServe(s.addressBuilder(), s.router)
	return nil
}

func (s *HttpServer) RegisterRouter(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	if s.router == nil {
		panic(NilRouterError)
	}
	s.router.HandleFunc(path, handler)
}

func (s *HttpServer) GroupRouter(path string, fn func(r *http.ServeMux)) {
	r := http.NewServeMux()
	fn(r)
	s.RegisterPathPrefix(path, r)
}

func (s *HttpServer) RegisterPathPrefix(path string, handler http.Handler) {
	s.router.Handle(path, handler)
}
