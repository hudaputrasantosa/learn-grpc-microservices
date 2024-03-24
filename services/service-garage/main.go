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

var localStorage *pb.GarageListByUser

type GaragesServer struct {
	pb.UnimplementedGarageServer
}

func main(){

}