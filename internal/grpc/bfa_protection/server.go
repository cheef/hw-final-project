package bfagrpc

import (
	"context"
	pb "github.com/cheef/hw-final-project/pkg/server/grpc/api/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type BFAProtection interface {
	Authorize(ctx context.Context, login, password, ip string) (bool, error)
	FlushBucket(ctx context.Context, login, ip string) error
	BlacklistAdd(ctx context.Context, cidr string) error
	BlacklistRemove(ctx context.Context, cidr string) error
	WhitelistAdd(ctx context.Context, cidr string) error
	WhitelistRemove(ctx context.Context, cidr string) error
}

type serverAPI struct {
	pb.UnimplementedBFAProtectionServer
	repo BFAProtection
}

func Register(grpcServer *grpc.Server, repo BFAProtection) {
	pb.RegisterBFAProtectionServer(grpcServer, &serverAPI{repo: repo})
}

func (s *serverAPI) Authorize(ctx context.Context, req *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	result, err := s.repo.Authorize(ctx, req.GetLogin(), req.GetPassword(), req.GetIp())

	return &pb.AuthorizeResponse{Ok: result}, err
}

func (s *serverAPI) FlushBucket(ctx context.Context, req *pb.FlushBucketRequest) (*empty.Empty, error) {
	err := s.repo.FlushBucket(ctx, req.GetLogin(), req.GetIp())

	return &empty.Empty{}, err
}

func (s *serverAPI) BlacklistAdd(ctx context.Context, req *pb.CIDRRequest) (*empty.Empty, error) {
	err := s.repo.BlacklistAdd(ctx, req.GetCidr())

	return &empty.Empty{}, err
}

func (s *serverAPI) BlacklistRemove(ctx context.Context, req *pb.CIDRRequest) (*empty.Empty, error) {
	err := s.repo.BlacklistRemove(ctx, req.GetCidr())

	return &empty.Empty{}, err
}

func (s *serverAPI) WhitelistAdd(ctx context.Context, req *pb.CIDRRequest) (*empty.Empty, error) {
	err := s.repo.WhitelistAdd(ctx, req.GetCidr())

	return &empty.Empty{}, err
}

func (s *serverAPI) WhitelistRemove(ctx context.Context, req *pb.CIDRRequest) (*empty.Empty, error) {
	err := s.repo.WhitelistRemove(ctx, req.GetCidr())

	return &empty.Empty{}, err
}
