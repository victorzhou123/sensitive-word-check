package main

import (
	"context"

	pb "github.com/victorzhou123/sensitive-word-check/proto"
)

type server struct {
	pb.UnimplementedCheckServer

	checker Checker
}

func NewServer() (pb.CheckServer, error) {
	checker, err := newChecker()
	if err != nil {
		return nil, err
	}

	return &server{
		checker: checker,
	}, nil
}

func (s *server) Check(ctx context.Context, word *pb.Word) (*pb.BoolMsg, error) {
	return &pb.BoolMsg{
		Val: s.checker.Check(word.Word),
	}, nil
}
