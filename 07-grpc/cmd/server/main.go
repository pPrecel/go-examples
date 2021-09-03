package main

import (
	"context"
	"io"
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

func (s *service) ValueToStream(v *pb.Value, service pb.HelloService_ValueToStreamServer) error {
	logrus.Infof("received: {Val: %d, Name: \"%s\"}", v.Val, v.Name)
	for i := 0; i < 5; i++ {
		err := service.Send(v)
		if err != nil {
			logrus.Fatalf("failed to send: {Val: %d, Name: \"%s\"}", v.Val, v.Name)
		}
	}

	logrus.Info("all values sent")
	return nil
}

func (s *service) StreamToValue(svc pb.HelloService_StreamToValueServer) error {
	logrus.Info("receving values from the stream")

	for {
		res, err := svc.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			logrus.Errorf("streaming error: %v", err)
			return err
		}
		logrus.Infof("received value from the stream: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
	}

	val := &pb.Value{
		Val:  66,
		Name: "calling back",
	}
	err := svc.SendAndClose(val)
	if err != nil {
		logrus.Fatalf("failed to send: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	}
	logrus.Infof("sending back: {Val: %d, Name: \"%s\"}", val.Val, val.Name)

	logrus.Info("ending connection")
	return nil
}

func (s *service) StreamToStream(svc pb.HelloService_StreamToStreamServer) error {
	logrus.Info("receving and sending values from the stream")

	for {
		res, err := svc.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			logrus.Errorf("streaming error: %v", err)
			return err
		}
		logrus.Infof("received value from the stream: {Val: %d, Name: \"%s\"}", res.Val, res.Name)

		err = svc.Send(res)
		if err != nil {
			logrus.Fatalf("failed to send: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
		}
		logrus.Infof("sent value to the stream: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
	}

	logrus.Info("ending connection")
	return nil
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
