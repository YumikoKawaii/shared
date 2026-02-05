package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Config struct {
	GRPC string `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP string `json:"http" mapstructure:"http" yaml:"http"`
}

func DefaultConfig() *Config {
	return &Config{
		GRPC: "0.0.0.0:10443",
		HTTP: "0.0.0.0:10080",
	}
}

type Server struct {
	instance *grpc.Server
	mux      *runtime.ServeMux
	config   *Config
}

func Initialize(config *Config, opts ...grpc.ServerOption) *Server {
	return &Server{
		instance: grpc.NewServer(opts...),
		mux:      runtime.NewServeMux(),
		config:   config,
	}
}

func (s *Server) Instance() *grpc.Server {
	return s.instance
}

func (s *Server) Mux() *runtime.ServeMux {
	return s.mux
}

func (s *Server) GRPCHost() string {
	return s.config.GRPC
}

func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	httpMux := http.NewServeMux()
	httpMux.Handle("/", s.mux)
	httpServer := http.Server{
		Addr:    s.config.HTTP,
		Handler: httpMux,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			errch <- err
		}
	}()
	go func() {
		listener, err := net.Listen("tcp", s.config.GRPC)
		if err != nil {
			errch <- err
			return
		}
		if err := s.instance.Serve(listener); err != nil {
			errch <- err
		}
	}()
	for {
		select {
		case <-stop:
			ctx := context.Background()
			httpServer.Shutdown(ctx)
			s.instance.GracefulStop()
			return nil
		case err := <-errch:
			return err
		}
	}
}
