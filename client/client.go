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

const (
	IP = 123
)

func Calculate(ctx context.Context, conn *grpc.ClientConn, num1 int, num2 int, oper string) (*pb.OpResponse, error) {

	c := pb.NewCalculatorClient(conn)
	req := &pb.OpRequest{Num1: int32(num1), Num2: int32(num2)}

	switch oper {

	case "+":
		return c.Add(ctx, req)
	case "-":
		return c.Sub(ctx, req)
	case "*":
		return c.Mul(ctx, req)
	case "/":
		return c.Div(ctx, req)

	default:
		return nil, fmt.Errorf("Invalid operation")
	}
}

func GetInputFromUser() (int, int, string, error) {
	var num1, num2 int
	var oper string

	fmt.Print("Type the first num: ")
	fmt.Scanln(&num1)
	fmt.Print("Type the operation ( + - / *): ")
	fmt.Scanln(&oper)
	fmt.Print("Type the second num: ")
	fmt.Scanln(&num2)

	return num1, num2, oper, nil
}

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Connection Failed")
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		num1, num2, oper, err := GetInputFromUser()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		res, err := Calculate(ctx, conn, num1, num2, oper)
		if err != nil {
			fmt.Printf("Calculate failed! %v\n", err)
		} else {
			fmt.Printf("result is: %v\n", res.Result)
		}
	}

}
