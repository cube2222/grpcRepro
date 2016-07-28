package main

import (
	"google.golang.org/grpc"
	"github.com/cube2222/grpcRepro/message"
	"golang.org/x/net/context"
	"fmt"
)

func main() {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	client := message.NewMyServiceClient(conn)
	res, _ := client.GetResponse(context.Background(), &message.Member{"localhost:3000", 15, 13, true})
	for _, item := range res.List {
		fmt.Printf("%v : %v : %v : %v", item.Address, item.Heartbeat, item.Heartbeat, item.Alive)
	}
	//time.Sleep(time.Second * 4)
	conn.Close()
}