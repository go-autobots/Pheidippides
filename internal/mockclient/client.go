/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 19:30 12月
 **/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	v1 "Pheidippides/api/pheidiqueue/v1"
)

func newClient(targetRPCAddr string) v1.PheidiQueueClient {
	conn, dialErr := grpc.Dial(targetRPCAddr,
		grpc.WithInsecure(), grpc.WithBlock(),
	)
	if dialErr != nil {
		panic(dialErr)
	}
	defer conn.Close()
	return v1.NewPheidiQueueClient(conn)
}

func main() {
	target := "localhost:9091"
	c := newClient(target)
	mockSend(c)
	//mockGetFrom(c)
	//mockSend(c)
}

func mockGetFrom(c v1.PheidiQueueClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, getErr := c.GetFrom(ctx, &v1.Pheidippides{Topic: "qqqq"})
	if getErr != nil {
		fmt.Println("get err", getErr)
	}

	fmt.Println("response content ---->", r.GetContent())
}

func mockSend(c v1.PheidiQueueClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, sendErr := c.SendTo(ctx, &v1.Pheidippides{Topic: "qqqq", Content: "123"})
	if sendErr != nil {
		fmt.Println("send err", sendErr)
	}

	fmt.Println("send  res ---->", r.GetIsSuccess())
}
