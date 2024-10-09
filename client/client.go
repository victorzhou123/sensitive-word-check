package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/victorzhou123/sensitive-word-check/proto"
)

type SensitiveWordChecker interface {
	Check(string) bool
}

type sensitiveWordCheckClient struct {
	conn   *grpc.ClientConn
	expire time.Duration

	client pb.CheckClient
}

func NewSensitiveWordCheckClient(addr string, t time.Duration) (SensitiveWordChecker, error) {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &sensitiveWordCheckClient{
		conn:   conn,
		expire: t,
		client: pb.NewCheckClient(conn),
	}, nil
}

func (c *sensitiveWordCheckClient) Check(word string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), c.expire)
	defer cancel()

	val, err := c.client.Check(ctx, &pb.Word{Word: word})
	if err != nil {
		log.Fatalf("check error: %s", err.Error())

		return false
	}

	return val.Val
}
