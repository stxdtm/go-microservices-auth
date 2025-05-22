package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	gp "github.com/stxdtm/go-microservices-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	gp.UnimplementedUserV1Server
}

// переопределяем метод Get, чтобы вместо заглушки использовать нужный функционал
func (s *server) Get(ctx context.Context, req *gp.GetRequest) (*gp.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &gp.GetResponse{
		User: &gp.User{
			Id: req.GetId(),
			Info: &gp.UserInfo{
				Name:  gofakeit.BeerName(),
				Email: gofakeit.Email(),
				Role:  1,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	gp.RegisterUserV1Server(s, &server{})
	log.Printf("server listening on %d\n", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatal("listenint error: ", err)
	}

}
