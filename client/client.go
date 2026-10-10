package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ronmandeles/gRPC-Calculator/pkg/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Connection Failed")
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	var num1, num2 int
	var oper string

	fmt.Print("Type the first num: ")
	fmt.Scanln(&num1)
	fmt.Print("Type the operation ( + - / *): ")
	fmt.Scanln(&oper)
	fmt.Print("Type the second num: ")
	fmt.Scanln(&num2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	switch {
	case oper == "+":
		res, err := c.Add(ctx, &pb.OpRequest{Num1: int32(num1), Num2: int32(num2), Operation: oper})
		if err != nil {
			log.Fatalf("failed to add, %v", err)
		}
		fmt.Printf("The result is %v", res)

	case oper == "-":
		res, err := c.Sub(ctx, &pb.OpRequest{Num1: int32(num1), Num2: int32(num2), Operation: oper})
		if err != nil {
			log.Fatalf("failed to sub %v", err)
		}
		fmt.Printf("The result is %v", res)

	case oper == "*":
		res, err := c.Mul(ctx, &pb.OpRequest{Num1: int32(num1), Num2: int32(num2), Operation: oper})
		if err != nil {
			log.Fatalf("failed to mul %v", err)
		}
		fmt.Printf("The result is %v", res)

	case oper == "/":
		res, err := c.Div(ctx, &pb.OpRequest{Num1: int32(num1), Num2: int32(num2), Operation: oper})
		if err != nil {
			log.Fatalf("failed to div %v", err)
		}
		fmt.Printf("The result is %v", res)

	}

}
