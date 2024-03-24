package main

import (
	"context"
	"log"
	"net"

	"learn-grpc/common/config"
	"learn-grpc/common/pb"
	"learn-grpc/common/proto"

	"google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/grpc"
)

var localStorage *pb.UserList

type UsersServer struct{
	pb.UnimplementedUsersServer
}

func (UsersServer) Register(_ context.Context, param *pb.User) (*emptypb.Empty, error){
	user := param

	localStorage.List = append(localStorage.List, user)
	log.Println("Registring User", user.String())

	return new(emptypb.Empty), nil
}

func (UsersServer) List(context.Context, *emptypb.Empty) (*pb.UserList, error){
	return localStorage, nil
}

func init(){
	localStorage = new(pb.UserList)
	localStorage.List = make([]*pb.UserList, 0)
}

func main(){
	srv := grpc.NewServer()
	var userSrv UsersServer
	pb.RegisterUsersServer(srv, userSrv);
	
	log.Println("Starting Grpc Server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
if err != nil {
    log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
}

log.Fatal(srv.Serve(l))
}