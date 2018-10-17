package main

import (
	"log"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb"merlot_project/merlot_write/service"

)

const (
	address     = "172.16.3.168:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArticleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	action := &pb.Action{ActionValue:pb.ActionType_ARTICLES_QUERY}
	var articles = []*pb.Article{&pb.Article{ArticleID:1,ArticleText:"天气不错"}}
	in := pb.ArticleRequest{Action:action,Article:articles}
	r, err := c.RouteArticle(ctx,&in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("client: %s", r.Message)
}