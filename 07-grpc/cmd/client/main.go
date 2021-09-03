package main

import (
	"context"
	"io"
	"time"

	pb "github.com/pPrecel/go-examples/07-grpc/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// getValues gets values from the server via stream
func getValues(c pb.HelloServiceClient, val *pb.Value) {
	logrus.Info("sending request: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	svc, err := c.ValueToStream(ctx, val)
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}

	for {
		res, err := svc.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			logrus.Fatalf("streaming error: %v", err)
		}

		logrus.Infof("received value from the stream: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
	}

	logrus.Info("ending connection")
}

// getValue gets value from the server
func getValue(c pb.HelloServiceClient, val *pb.Value) {
	logrus.Infof("sending request: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := c.GetValue(ctx, val)
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}
	logrus.Infof("received value: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
}

// streamValues sends a few values and then receive and log one value
func streamValues(c pb.HelloServiceClient, val *pb.Value) {
	logrus.Infof("sending 5x request: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	svc, err := c.StreamToValue(ctx)
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}

	for i := 0; i < 5; i++ {
		err := svc.Send(val)
		if err != nil {
			logrus.Fatalf("streaming error: %v", err)
		}
		logrus.Infof("value sent: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	}

	res, err := svc.CloseAndRecv()
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}
	logrus.Infof("received value: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
}

// chat chats with the server :)
func chat(c pb.HelloServiceClient, val *pb.Value) {
	logrus.Infof("receiving and sending 5x request: {Val: %d, Name: \"%s\"}", val.Val, val.Name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	svc, err := c.StreamToStream(ctx)
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}

	for i := 0; i < 5; i++ {
		err := svc.Send(val)
		if err != nil {
			logrus.Fatalf("streaming error: %v", err)
		}
		logrus.Infof("value sent: {Val: %d, Name: \"%s\"}", val.Val, val.Name)

		res, err := svc.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			logrus.Fatalf("streaming error: %v", err)
		}
		logrus.Infof("received value from the stream: {Val: %d, Name: \"%s\"}", res.Val, res.Name)
	}

	err = svc.CloseSend()
	if err != nil {
		logrus.Fatalf("can't get value from the server: %v", err)
	}

	logrus.Info("chat ended")
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	getValue(client, &pb.Value{Val: 4, Name: "hello test name"})

	// getValues(client, &pb.Value{Val: 7, Name: "hello another test"})

	// streamValues(client, &pb.Value{Val: 10, Name: "hello another another test"})

	// chat(client, &pb.Value{Val: 20, Name: "hello chatting machine"})
}
