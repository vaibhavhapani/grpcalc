package main

import (
	"context"
	//"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/vaibhavhapani/grpcalc.git/pb"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
)

func main(){
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address in the format of host:port",
	)
	flag.Parse()

	//creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	// opts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(creds),
	// }

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}	

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil{
		log.Fatalln("Fail to dial:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 10, 10, 10},
	})
	if err != nil{
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("result:", res.Result)
}