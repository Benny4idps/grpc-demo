package main

import (
	"context"
	pb "grpc-demo/blog/proto"
	"log"
)

func  readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error was happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}