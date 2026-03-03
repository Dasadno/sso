package auth

import (
	"context"
	"log/slog"

	ssov1 "github.com/Dasadno/sso-protos/gen/go/proto/sso"
	"github.com/Dasadno/sso/internal/grpc/auth/models"
	"github.com/Dasadno/sso/internal/grpc/validate"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponce, error) {
	l := models.Login{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	err := validate.V.Struct(l)
	if err != nil {
		slog.Error("Validation error")
		return nil, status.Error(codes.InvalidArgument, "Error: Invalid argument")
	}
	return &ssov1.LoginResponce{
		Token: uuid.NewString(), //Generate JWT must be here
	}, nil
}

func (s *serverApi) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponce, error) {
	l := models.Register{
		Email:             req.GetEmail(),
		Password:          req.GetPassword(),
		PasswordConfirmed: req.GetPasswordConfirmed(),
	}
	err := validate.V.Struct(l)
	if err != nil {
		slog.Error("Validation error")
		return nil, status.Error(codes.InvalidArgument, "Error: invalid argument")
	}
	return &ssov1.RegisterResponce{
		Id: int64(uuid.Max.ID()),
	}, nil
}
func (s *serverApi) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponce, error) {
	panic("impolement me")
}
