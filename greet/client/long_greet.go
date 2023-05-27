package main

import (
	"context"
	pb "grpc-demo/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient)  {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest {
		{ FirstName: "Benny" },
		{ FirstName: "Helena" },
		{ FirstName: "Test" },
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling long greet%v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving response from longGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

