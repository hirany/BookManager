package main

import (
	pb "../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func client(sn, bn int64, sw bool) string {

	var title string
	var err error
	connection, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewLendServiceClient(connection)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Lend(context, &pb.LendRequest{BookNumber: bn, StudentNumber: sn, Sw: sw})
	var str string
	if sw {
		str = "Rental "
	} else {
		str = "Return "
	}
	if err != nil {
		fmt.Println(str + " was failed")
		log.Fatal(err)
	} else {
		fmt.Println(str + "was successful")
		fmt.Println("student number : ", response.GetStudentNumber())
		title, err = janToTitle(bn)
		if err != nil {
			log.Fatal(err)
		}
	}
	return title
}
