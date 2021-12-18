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

func main() {
	target := "localhost:9091"
	conn, dialErr := grpc.Dial(target,
		grpc.WithInsecure(), grpc.WithBlock(),
	)
	if dialErr != nil {
		panic(dialErr)
	}
	defer conn.Close()
	c := v1.NewPheidiQueueClient(conn)
	mockGetFrom(c)
	mockSend(c)
}

func mockGetFrom(c v1.PheidiQueueClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, getErr := c.GetFrom(ctx, &v1.Pheidippides{Topic: "xxxxx"})
	if getErr != nil {
		panic(getErr)
	}

	fmt.Println("response content ---->", r.GetContent())
}

func mockSend(c v1.PheidiQueueClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, sendErr := c.SendTo(ctx, &v1.Pheidippides{Topic: "qqqq", Contents: []string{"1", "2"}})
	if sendErr != nil {
		panic(sendErr)
	}

	fmt.Println("send  res ---->", r.GetIsSuccess())
}
