package main

import (
	"context"
	pb "grpc-demo/blog/proto"
	"log"
)
func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("Update blog was invoked")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Not Benny",
		Title: "New title",
		Content: "Contennnnnnnt",
	}

	_, err := c.Update(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}