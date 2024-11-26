package user

import (
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"google.golang.org/grpc"
)

func (u *UsersGrpc) Register(grpcServer *grpc.Server) {
	grpc_gen.RegisterUserServiceServer(grpcServer, u)
}
