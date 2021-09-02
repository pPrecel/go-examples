package main

import (
	"context"
	"net"

	pb "github.com/pPrecel/go-examples/07-grpc/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedHelloServiceServer
}

func (s *service) GetValue(c context.Context, v *pb.Value) (*pb.Value, error) {
	logrus.Infof("received: {Val: %d, Name: \"%s\"}", v.Val, v.Name)
	logrus.Infof("sending back: {Val: %d, Name: \"%s\"}", v.Val, v.Name)
	return v, nil
}

func main() {
	logrus.Info("Starting server...")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	pb.RegisterHelloServiceServer(grpcServer, &service{})

	logrus.Info("Listen And Serve...")
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
