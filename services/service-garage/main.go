// file ini merupakan simulasi aplikasi microservice garage yang akan saling terhubung dengan aplikasi main

package main

import (
	"context"
	"log"
	"net"

	"learn-grpc/common/config"
	"learn-grpc/common/pb"

	"google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/grpc"
)

var localStorage *pb.GarageListByUser

func init(){
	localStorage = new(pb.GarageListByUser)
	localStorage.List = make(map[string]*pb.GarageList)
}

type GaragesServer struct {
	pb.UnimplementedGaragesServer
}

// menjelaskan mengenai logic Add pada service user proto
func (GaragesServer) Add(_ context.Context, param *pb.GarageAndUserId)(*emptypb.Empty, error){
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(pb.GarageList)
		localStorage.List[userId].List = make([]*pb.Garage, 0)
	}

	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)
	 log.Println("Adding garage", garage.String(), "for user", userId)

   return new(emptypb.Empty), nil
}

//menjelaska  megenai logic List pada service user proto
func (GaragesServer) List(_ context.Context, param *pb.GarageUserId) (*pb.GarageList, error) {
   userId := param.UserId

   return localStorage.List[userId], nil
}

// menjalankan fungsi utama dan server aplikasi
func main(){
	srv := grpc.NewServer()
    var garageSrv GaragesServer
    pb.RegisterGaragesServer(srv, garageSrv)

    log.Println("Starting RPC server at", config.ServiceGaragePort)

    l, err := net.Listen("tcp", config.ServiceGaragePort)
    if err != nil {
        log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
    }

    log.Fatal(srv.Serve(l))

}