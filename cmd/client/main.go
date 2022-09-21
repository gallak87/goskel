package main

import (
	"context"
	"log"
	"time"

	"github.com/gallak87/goskel/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:9090", creds, grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := userpb.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &userpb.GetUserRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("UserClient: server response:\n\t%s\n\n", r.Name)
}
