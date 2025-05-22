package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fatih/color"
	gp "github.com/stxdtm/go-microservices-auth/pkg/user_v1"
)

const (
	grpcServer = "localhost:50051"
	userID     = 100
)

func main() {
	conn, err := grpc.NewClient(grpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := gp.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &gp.GetRequest{Id: userID})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(color.RedString("User request:\n"), color.GreenString(":: %+v\n", r.GetUser()))

}
