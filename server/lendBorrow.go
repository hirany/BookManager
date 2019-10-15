package main

import (
	"context"
	"errors"
	"fmt"

	pb "../pb"
)

type lendMemory struct{}

func (s *lendMemory) Lend(ctx context.Context, req *pb.LendRequest) (*pb.LendResponse, error) {
	sn := userTable{id: req.GetStudentNumber()}
	bn := bookTable{id: req.GetBookNumber()}
	if bn.id == 0 || sn.id == 0 {
		return &pb.LendResponse{BookNumber: bn.id, StudentNumber: sn.id}, errors.New("BookNumber value or studentNumber value is 0")
	}
	exist := sn.existUserData()
	if !exist {
		return &pb.LendResponse{BookNumber: bn.id, StudentNumber: sn.id}, errors.New("Not found: user")
	}
	exist = bn.existBookData()
	if !exist {
		bn.insertBookData()
	}
	insertReceiptData(bn, sn)
	fmt.Println("Lending information")
	fmt.Println("student number :", sn)
	fmt.Println("book number    :", bn)
	return &pb.LendResponse{BookNumber: bn.id, StudentNumber: sn.id}, nil
}
