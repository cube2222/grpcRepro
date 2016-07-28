package main

import (
	"github.com/cube2222/grpcRepro/message"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", "localhost:3000")
	s := grpc.NewServer()
	myServer := MyServer{}
	message.RegisterMyServiceServer(s, myServer)
	s.Serve(lis)
}

type MyServer struct {
}

func (s MyServer) GetResponse(ctx context.Context, in *message.Member) (*message.MemberList, error) {
	members := &message.MemberList{List: make(map[string]*message.Member)}

	members.List[in.Address] = in

	members.List["0.0.0.0"] = &message.Member{Address:"0.0.0.0", Heartbeat:10, Timestamp:10, Alive:true}

	return members, nil
}
