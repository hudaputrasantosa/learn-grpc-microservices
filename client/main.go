package main

import (
	"fmt"
	"log"
	"context"
	    "strings"
    "encoding/json"
	

	"learn-grpc/common/config"
	"learn-grpc/common/pb"
		"google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

)

func serviceGarage() pb.GaragesClient {
  port := config.ServiceGaragePort
  conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
      log.Fatal("could not connect to", port, err)
  }

  return pb.NewGaragesClient(conn)
}

func serviceUser() pb.UsersClient {
  port := config.ServiceUserPort
  conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
      log.Fatal("could not connect to", port, err)
  }

  return pb.NewUsersClient(conn)
}

func main(){

	var user1 = pb.User{
		Id: "uu-45",
		Name: "Huda",
		Password: "huda-123",
		Gender: pb.UserGender_MALE,

	}

// 	var userList = &pb.UserList{
//      List: []*pb.User{
//          user1,
//      },
//  }

 var garage1 = pb.Garage{
     Id:   "g001",
     Name: "Kalimdor",
     Coordinate: &pb.GarageCoordinate{
         Latitude:  23.2212847,
         Longitude: 53.22033123,
     },
 }

 var garage2 = pb.Garage{
     Id:   "g002",
     Name: "Karimun",
     Coordinate: &pb.GarageCoordinate{
         Latitude:  23.2212847,
         Longitude: 53.22033123,
     },
 }

//  var garageList = &pb.GarageList{
//      List: []*pb.Garage{
//          garage1,
//      },
//  }

//  var garageListByUser = &pb.GarageListByUser{
//      List: map[string]*pb.GarageList{
//          user1.Id: garageList,
//      },
//  }

// =========== as json string
// jsonb, err1 := protojson.Marshal(garageList)
// if err1 != nil {
//    fmt.Println(err1.Error())
//    os.Exit(0)
// }

user := serviceUser()
fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))
user.Register(context.Background(), &user1)

// show all registered users
res1, err := user.List(context.Background(), new(emptypb.Empty))
if err != nil {
    log.Fatal(err.Error())
}
res1String, _ := json.Marshal(res1.List)
log.Println(string(res1String))


garage := serviceGarage()

fmt.Printf("\n %s> garage test A\n", strings.Repeat("=", 10))

// add garage1 to user1
garage.Add(context.Background(), &pb.GarageAndUserId{
   UserId: user1.Id,
   Garage: &garage1,
})

// add garage2 to user1
garage.Add(context.Background(), &pb.GarageAndUserId{
   UserId: user1.Id,
   Garage: &garage2,
})

// show all garages of user1
res2, err := garage.List(context.Background(), &pb.GarageUserId{UserId: user1.Id})
if err != nil {
   log.Fatal(err.Error())
}
res2String, _ := json.Marshal(res2.List)
log.Println(string(res2String))

}