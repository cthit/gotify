package grpc

import (
	"context"
	"fmt"
	"github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	rpcPort, webPort string
	debug            bool
	mailService      mail.Service
	wg               sync.WaitGroup
}

func NewServer(rpcPort, webPort string, debug bool, mailService mail.Service) (*Server, error) {
	return &Server{
		rpcPort:     rpcPort,
		webPort:     webPort,
		debug:       debug,
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
	grpcServer.Serve(lis)
	return nil
}

func (s *Server) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gotify.RegisterMailerHandlerFromEndpoint(ctx, mux, "localhost:"+s.rpcPort, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":"+s.webPort, mux)
}
