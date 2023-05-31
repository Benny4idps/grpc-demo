package main

import (
	"context"
	pb "grpc-demo/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Create blog was invoked")

	blog := &pb.Blog{
		AuthorId: "Benny",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)

	return res.Id
}
