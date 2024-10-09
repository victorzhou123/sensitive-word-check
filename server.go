package main

import (
	"context"

	pb "github.com/victorzhou123/sensitive-word-check/proto"
)

type server struct {
	pb.UnimplementedCheckServer

	checker Checker
}

func NewServer() pb.CheckServer {
	return &server{
		checker: newChecker(),
	}
}

func (s *server) Check(ctx context.Context, word *pb.Word) (*pb.BoolMsg, error) {
	return &pb.BoolMsg{
		Val: s.checker.Check(word.Word),
	}, nil
}
