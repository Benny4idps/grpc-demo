package main

import (
	"context"
	pb "grpc-demo/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Println("doGreetEveryOne was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{ FirstName: "Benny" },
		{ FirstName: "Helena" },
		{ FirstName: "Test" },
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
	}()

	go func()  {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiveing: %v\n", err)
				break
			}

			log.Printf("Recieved: %v\n", res.Result)
		}
		
		close(waitc)
	}()

	<-waitc
}


