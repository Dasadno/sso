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
		Email:    req.Email,
		Password: req.Password,
	}
	err := validate.V.Struct(l)
	if err != nil {
		slog.Error("Incorrect enter")
		return nil, status.Error(codes.InvalidArgument, "Error: Invalid argument")
	}
	return &ssov1.LoginResponce{
		Token: uuid.NewString(),
	}, nil
}

func (s *serverApi) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponce, error) {
	panic("implement me")
}
func (s *serverApi) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponce, error) {
	panic("impolement me")
}
