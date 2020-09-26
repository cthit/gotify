package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/cthit/gotify/internal/app/config"
	gotify "github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
)

type Server struct {
	rpcPort, webPort string
	debug            bool
	env              string
	mailService      mail.Service
	wg               sync.WaitGroup
}

func NewServer(rpcPort, webPort string, env string, debug bool, mailService mail.Service) (*Server, error) {
	return &Server{
		rpcPort:     rpcPort,
		webPort:     webPort,
		debug:       debug,
		env:         env,
		mailService: mailService,
	}, nil
}

func (s *Server) Start() {
	s.wg.Add(1)

	go func() {
		err := s.startGRPC()
		fmt.Println(err)
		s.wg.Done()
	}()

	s.wg.Add(1)

	go func() {
		err := s.startREST()
		fmt.Println(err)
		s.wg.Done()
	}()

	s.wg.Wait()
}

func (s *Server) startGRPC() error {
	lis, err := net.Listen("tcp", ":"+s.rpcPort)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	gotify.RegisterMailerServer(grpcServer, s)

	return grpcServer.Serve(lis)
}

func (s *Server) startREST() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mux http.Handler = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gotify.RegisterMailerHandlerFromEndpoint(ctx, mux.(*runtime.ServeMux), "localhost:"+s.rpcPort, opts)
	if err != nil {
		return err
	}

	if s.env == config.EnvDevelopment {
		mux = allowCORS(mux)
	}

	return http.ListenAndServe(":"+s.webPort, mux)
}