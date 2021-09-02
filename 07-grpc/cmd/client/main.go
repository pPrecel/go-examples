package main

import (
	"context"

	pb "github.com/pPrecel/go-examples/07-grpc/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getValue(c pb.HelloServiceClient) {
	val := &pb.Value{
		Val:  4,
		Name: "hello test name",
	}
	logrus.Infof("sending request: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	val, err := c.GetValue(context.Background(), val)
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}

	logrus.Infof("received value: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	getValue(client)
}
