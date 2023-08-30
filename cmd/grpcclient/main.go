package main

import (
	"context"
	"math/rand"
	"time"

	entpb "github.com/ecshreve/dndgen/ent/proto/entpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a User service Client on the connection.
	client := entpb.NewAbilityScoreServiceClient(conn)

	// Ask the server to create a random User.
	ctx := context.Background()

	created, err := client.Create(ctx, &entpb.CreateAbilityScoreRequest{
		AbilityScore: &entpb.AbilityScore{
			Indx:     "test_indx4",
			Name:     "Test Indx4",
			FullName: "Testivus Indexius4",
			Desc:     []string{"test2"},
		},
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating as: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := client.Get(ctx, &entpb.GetAbilityScoreRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving as: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved as with id=%d: %v", get.Id, get)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err = client.Get(ctx, &entpb.GetAbilityScoreRequest{
		Id: 6,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving as: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved as with id=%d: %v", get.Id, get)
}
