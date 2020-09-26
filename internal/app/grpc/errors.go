package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcError struct {
	error
	status codes.Code
}

func (ge grpcError) GRPCStatus() *status.Status {
	return status.New(ge.status, ge.Error())
}

func WithErrorStatus(err error, status codes.Code) grpcError {
	return grpcError{
		error:  err,
		status: status,
	}
}