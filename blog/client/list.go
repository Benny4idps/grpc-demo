package main

import (
	"context"
	pb "grpc-demo/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient)  {
	log.Println("ListBlog was invoked")

	stream, err := c.ListBlogd(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling list blogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happned: %v\n", err)
		}

		log.Println(res)
	}
}