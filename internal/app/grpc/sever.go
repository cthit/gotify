package grpc

import (
	"context"
	"github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	rpcPort, webPort string
	debug            bool
	mailService      mail.MailService
	wg               sync.WaitGroup
}

func (s *Server) SendMail(_ context.Context, r *gotify.SendMailRequest) (*gotify.SendMailResponse, error) {
	//validate yo
	m, err := s.mailService.SendMail(mail.Mail{
		To:      r.Mail.To,
		From:    r.Mail.From,
		Subject: r.Mail.Subject,
		Body:    r.Mail.Body,
	})
	if err != nil {
		return nil, err
	}
	//handle them errors
	return &gotify.SendMailResponse{
		Mail: &gotify.Mail{
			To:      m.To,
			From:    m.From,
			Subject: m.Subject,
			Body:    m.Body,
		},
		Success: true,
		Error:   "",
	}, nil
}

func NewServer(rpcPort, webPort string, debug bool, mailService mail.MailService) (*Server, error) {
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
		log.Fatal(s.startGRPC())
		s.wg.Done()
	}()
	s.wg.Add(1)
	go func() {
		log.Fatal(s.startREST())
		s.wg.Done()
	}()
	s.wg.Wait()
}

func (s *Server) startGRPC() error {
	lis, err := net.Listen("tcp", ":" + s.rpcPort)
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
	err := gotify.RegisterMailerHandlerFromEndpoint(ctx, mux, ":" + s.rpcPort, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":" + s.webPort, mux)
}
