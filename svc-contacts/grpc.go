package main

import (
	"context"

	"github.com/bweston92/aip-158-demo/svc-contacts/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	grpcServer struct{}
)

var _ (api.ContactsServer) = (*grpcServer)(nil)

func (s *grpcServer) ListContacts(context.Context, *api.ListContactsRequest) (*api.ListContactsResponse, error) {
	return nil, status.New(codes.Unimplemented, "not implemented").Err()
}
