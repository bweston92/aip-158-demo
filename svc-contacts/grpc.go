package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/bweston92/aip-158-demo/svc-contacts/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	grpcServer struct {
		contacts Store
	}
)

var _ (api.ContactsServer) = (*grpcServer)(nil)

func startPublicAPI(contacts Store) error {
	nl, err := net.Listen("tcp", ":3101")
	if err != nil {
		return fmt.Errorf("unable to open port 3101: %w", err)
	}

	srv := grpc.NewServer()
	api.RegisterContactsServer(srv, &grpcServer{
		contacts: contacts,
	})
	go srv.Serve(nl)

	return nil
}

func validateRequest(req *api.ListContactsRequest) error {
	if req.PageSize < 0 || req.PageSize > 30 {
		return errors.New("specify a page size between 0 and 30")
	} else if req.PageSize == 0 {
		req.PageSize = 10
	}

	return nil
}

func (s *grpcServer) ListContacts(ctx context.Context, req *api.ListContactsRequest) (*api.ListContactsResponse, error) {
	if err := validateRequest(req); err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	total, err := s.contacts.Count(ctx, Filters{})
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	items, err := s.contacts.List(ctx, Filters{}, 0, req.PageSize)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	resp := make([]*api.Contact, len(items))
	for i, v := range items {
		resp[i] = &api.Contact{
			Id:          v.ID,
			Forename:    v.Forename,
			Surname:     v.Surname,
			PhoneNumber: v.PhoneNumber,
		}
	}

	return &api.ListContactsResponse{
		Contacts:      resp,
		Total:         total,
		NextPageToken: "abcd",
	}, nil
}
