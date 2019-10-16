package main

import (
	pb "../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	for {
		sw, sn, bn := input()

		fmt.Println(sw)
		connection, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close()

		client := pb.NewLendServiceClient(connection)

		context, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := client.Lend(context, &pb.LendRequest{BookNumber: bn, StudentNumber: sn})
		if err != nil {
			fmt.Println("NG")
			log.Fatal(err)
		} else {
			fmt.Println("OK")
			fmt.Println("book number    : ", response.GetBookNumber())
			fmt.Println("student number : ", response.GetStudentNumber())
		}
	}
}

func input() (sw, sn, bn int64) {
	fmt.Println("lend: 1, borrow: 2")
	fmt.Scan(&sw)
	fmt.Print("student number: ")
	fmt.Scan(&sn)
	fmt.Print("book number: ")
	fmt.Scan(&bn)
	return
}
