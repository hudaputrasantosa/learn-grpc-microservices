package main

import (
	"fmt"
	// "os"
	"learn-grpc/common/pb"
)

func main(){

	var user1 = &pb.User{
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

//  var garage1 = &pb.Garage{
//      Id:   "g001",
//      Name: "Kalimdor",
//      Coordinate: &pb.GarageCoordinate{
//          Latitude:  23.2212847,
//          Longitude: 53.22033123,
//      },
//  }

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

// fmt.Printf("# ==== As JSON String\n       %s \n", string(jsonb))

	fmt.Printf("#========= 	Original \n	%#v \n", user1.String())

}